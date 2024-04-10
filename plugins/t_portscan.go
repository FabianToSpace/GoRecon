package plugins

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"slices"
	"strings"

	"github.com/FabianToSpace/GoRecon/config"
)

type PortScan struct {
	Name             string   // Name of the Portscan
	Description      string   // Description of the Portscan
	Type             string   // Type (e.g. TCP or UPD)
	Tags             []string // ['default', 'default-portscan']
	Arguments        []string // Command Arguments to add
	Command          string   // Executeable Command
	TargetAppend     bool     // Append Target String to the end of the Command
	OutputFormat     string
	ArgumentsInPlace bool
}

func (p PortScan) ReplaceInArguments(token, value string) []string {
	for i, arg := range p.Arguments {
		p.Arguments[i] = strings.Replace(arg, token, value, -1)
	}
	return p.Arguments
}

func (p PortScan) TokenizeOutput(target string) string {
	curDir, _ := os.Getwd()
	outputString := strings.Replace(p.OutputFormat, "{{.Target}}", target, -1)
	outputString = curDir + "/" + outputString
	return outputString
}

func (p PortScan) TokenizeArguments(target string) []string {
	if p.OutputFormat != "" {
		outputString := p.TokenizeOutput(target)
		p.ReplaceInArguments("{{.OutputFile}}", outputString)
	}

	if p.ArgumentsInPlace {
		p.Arguments = p.ReplaceInArguments("{{.PortRange}}", Config.PortRange)
	}

	if p.TargetAppend {
		p.Arguments = append(p.Arguments, target)
	} else {
		p.Arguments = append([]string{target}, p.Arguments...)
	}
	return p.Arguments
}

func (p PortScan) Run(target string) []Service {
	Init()
	if !slices.Contains(config.AllowedCommands, p.Command) {
		panic(fmt.Sprintf("Command %s is not allowed", p.Command))
	}

	Logger.Info(p.Name, target, "Starting"+p.Description)

	services := make(chan Service)

	args := p.TokenizeArguments(target)

	go p.executeCommand(target, args, services)

	var results []Service

	for r := range services {
		results = append(results, r)
	}

	<-services

	return results
}

func (p PortScan) executeCommand(target string, args []string, svc chan Service) {
	reader, writer := io.Pipe()

	cmdCtx, cmdDone := context.WithCancel(context.Background())

	cmd := exec.CommandContext(cmdCtx, p.Command, args...)
	cmd.Stdout = writer
	cmd.Stderr = writer

	go scanOutput(reader, p.Name, target, svc)

	_ = cmd.Start()
	go func() {
		_ = cmd.Wait()
		cmdDone()
		writer.Close()
		close(svc)
	}()

	<-cmdCtx.Done()
}

func scanOutput(reader *io.PipeReader, serviceName, target string, svc chan Service) {
	defer reader.Close()
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()

		Logger.Debug(serviceName, target, line)

		service := extractService(target, serviceName, line)
		if service != (Service{}) {
			svc <- service
		}
	}
}
