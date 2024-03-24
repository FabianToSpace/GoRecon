package plugins

import (
	"bufio"
	"context"
	"fmt"
	"gorecon/logger"
	"io"
	"os/exec"
	"regexp"
	"strconv"
)

func Whatweb() ServiceScan {
	moduleName := "whatweb"
	return ServiceScan{
		Name:        moduleName,
		Description: "Directory Buster",
		Tags:        []string{"default", "http"},
		Run: func(service Service) bool {
			match, _ := regexp.MatchString("^http", service.Name)
			if match {
				target := fmt.Sprintf("%s:%d", service.Target, service.Port)
				logger.Logger().Start(moduleName, target, "Starting WhatWeb")

				reader, writer := io.Pipe()

				cmdCtx, cmdDone := context.WithCancel(context.Background())

				scannerStopped := make(chan struct{})
				go func() {
					defer close(scannerStopped)

					scanner := bufio.NewScanner(reader)
					for scanner.Scan() {
						line := scanner.Text()
						logger.Logger().Debug(moduleName, service.Target, line)
					}
				}()

				cmd := exec.Command("whatweb", "--color=never", "--no-errors", "-a 3", "-v", service.Target+":"+strconv.Itoa(service.Port))

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

				logger.Logger().Done(moduleName, service.Target, "Done")
				return true
			}
			return false
		},
	}
}
