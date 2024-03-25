package plugins

import (
	"bufio"
	"context"
	"fmt"
	"gorecon/logger"
	"io"
	"os/exec"
	"regexp"
	"strings"
)

type ServiceScan struct {
	Name          string   // Name of the ServiceScan
	Description   string   // Description of the ServiceScan
	Type          string   // Type (e.g. TCP or UPD)
	Tags          []string // ['default', 'default-portscan']
	Arguments     []string // Command Arguments to add
	Command       string   // Executeable Command
	TargetAppend  bool     // Append Target String to the end of the Command
	TargetInplace bool
	TargetFormat  string
	MatchPattern  string
}

func (s ServiceScan) MatchCondition(service Service) bool {
	condition, _ := regexp.MatchString(s.MatchPattern, s.Name)

	return condition
}

func (s ServiceScan) TokenizeArguments(service Service) []string {
	targetString := strings.Replace(s.TargetFormat, "{{.Target}}", service.Target, -1)
	targetString = strings.Replace(targetString, "{{.Port}}", fmt.Sprintf("%d", service.Port), -1)
	targetString = strings.Replace(targetString, "{{.Scheme}}", service.Scheme, -1)

	if s.TargetInplace {
		// Replace the string {{.TargetPos}} in any of the array elements of Arguments
		for i, arg := range s.Arguments {
			s.Arguments[i] = strings.Replace(arg, "{{.TargetPos}}", targetString, -1)
		}

		return s.Arguments
	}
	if s.TargetAppend {
		s.Arguments = append(s.Arguments, targetString)
	} else {
		s.Arguments = append([]string{targetString}, s.Arguments...)
	}
	return s.Arguments
}

func (s ServiceScan) Run(service Service) bool {
	if s.MatchCondition(service) {
		target := fmt.Sprintf("%s:%d", service.Target, service.Port)
		logger.Logger().Start(s.Name, target, "Starting "+s.Description)

		reader, writer := io.Pipe()

		cmdCtx, cmdDone := context.WithCancel(context.Background())

		scannerStopped := make(chan struct{})
		go func() {
			defer close(scannerStopped)

			scanner := bufio.NewScanner(reader)
			for scanner.Scan() {
				line := scanner.Text()
				logger.Logger().Debug(s.Name, service.Target, line)
			}
		}()

		args := s.TokenizeArguments(service)

		cmd := exec.Command(s.Command, args...)

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

		logger.Logger().Done(s.Name, service.Target, "Done")
		return true
	}
	return false
}
