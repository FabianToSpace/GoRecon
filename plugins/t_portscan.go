package plugins

import (
	"bufio"
	"context"
	"fmt"
	"gorecon/config"
	"gorecon/logger"
	"io"
	"os"
	"os/exec"
	"slices"
	"strings"
)

type PortScan struct {
	Name         string   // Name of the Portscan
	Description  string   // Description of the Portscan
	Type         string   // Type (e.g. TCP or UPD)
	Tags         []string // ['default', 'default-portscan']
	Arguments    []string // Command Arguments to add
	Command      string   // Executeable Command
	TargetAppend bool     // Append Target String to the end of the Command
	OutputFormat string
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

	if p.TargetAppend {
		p.Arguments = append(p.Arguments, target)
	} else {
		p.Arguments = append([]string{target}, p.Arguments...)
	}
	return p.Arguments
}

func (p PortScan) Run(target string) []Service {
	if !slices.Contains(config.AllowedCommands, p.Command) {
		panic(fmt.Sprintf("Command %s is not allowed", p.Command))
	}

	logger.Logger().Info(p.Name, target, "Starting"+p.Description)

	services := make([]Service, 0)
	reader, writer := io.Pipe()

	cmdCtx, cmdDone := context.WithCancel(context.Background())

	scannerStopped := make(chan struct{})
	go func() {
		defer close(scannerStopped)

		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			line := scanner.Text()

			logger.Logger().Debug(p.Name, target, line)

			service := extractService(target, p.Name, line)
			if service != (Service{}) {
				services = append(services, service)
			}
		}
	}()

	args := p.TokenizeArguments(target)

	cmd := exec.CommandContext(cmdCtx, p.Command, args...)
	cmd.Stdout = writer
	cmd.Stderr = writer

	_ = cmd.Start()
	go func() {
		_ = cmd.Wait()
		cmdDone()
		writer.Close()
	}()
	<-cmdCtx.Done()

	<-scannerStopped

	return services
}
