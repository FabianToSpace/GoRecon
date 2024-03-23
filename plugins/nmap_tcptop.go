package plugins

import (
	"bufio"
	"context"
	"gorecon/config"
	"gorecon/logger"
	"io"
	"os/exec"
)

func NmapTcpTop() PortScan {
	moduleName := "nmap-tcp-top"
	return PortScan{
		Name:        moduleName,
		Description: "nmap TCP top",
		Type:        "tcp",
		Tags:        []string{"default", "default-portscan"},
		Run: func(target string) []Service {
			logger.Logger().Info(moduleName, target, "Starting Nmap TCP top scan")

			services := make([]Service, 0)
			reader, writer := io.Pipe()

			cmdCtx, cmdDone := context.WithCancel(context.Background())

			scannerStopped := make(chan struct{})
			go func() {
				defer close(scannerStopped)

				scanner := bufio.NewScanner(reader)
				for scanner.Scan() {
					line := scanner.Text()

					if config.GetConfig().Debug {
						logger.Logger().Debug(moduleName, target, line)
					}

					service := extractService(target, moduleName, line)
					if service != (Service{}) {
						services = append(services, service)
					}
				}
			}()

			cmd := exec.Command("nmap", target, "-sC", "-sV", "-vvvv")
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
		},
	}
}
