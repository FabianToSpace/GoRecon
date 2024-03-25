package plugins

import (
	"bufio"
	"context"
	"gorecon/logger"
	"io"
	"os/exec"
)

type PortScan struct {
	Name         string   // Name of the Portscan
	Description  string   // Description of the Portscan
	Type         string   // Type (e.g. TCP or UPD)
	Tags         []string // ['default', 'default-portscan']
	Arguments    []string // Command Arguments to add
	Command      string   // Executeable Command
	TargetAppend bool     // Append Target String to the end of the Command
}

func (p PortScan) TokenizeArguments(target string) []string {
	if p.TargetAppend {
		p.Arguments = append(p.Arguments, target)
	} else {
		p.Arguments = append([]string{target}, p.Arguments...)
	}
	return p.Arguments
}

func (p PortScan) Run(target string) []Service {
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
