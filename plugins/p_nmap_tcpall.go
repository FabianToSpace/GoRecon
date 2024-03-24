package plugins

import (
	"bufio"
	"context"
	"gorecon/config"
	"gorecon/logger"
	"io"
	"os/exec"
)

func NmapTcpAll() PortScan {
	moduleName := "nmap-tcp-all"
	return PortScan{
		Name:        moduleName,
		Description: "nmap TCP all",
		Type:        "tcp",
		Tags:        []string{"default", "default-portscan"},
		Run: func(target string) []Service {
			logger.Logger().Info(moduleName, target, "Starting Nmap TCP all scan")

			services := make([]Service, 0)
			reader, writer := io.Pipe()

			cmdCtx, cmdDone := context.WithCancel(context.Background())

			scannerStopped := make(chan struct{})
			go func() {
				defer close(scannerStopped)

				scanner := bufio.NewScanner(reader)
				for scanner.Scan() {
					line := scanner.Text()

					logger.Logger().Debug(moduleName, target, line)

					service := extractService(target, moduleName, line)
					if service != (Service{}) {
						services = append(services, service)
					}
				}
			}()

			cmd := exec.Command("nmap", target, "-sC", "-sV", "-p"+config.GetConfig().PortRange, "-vvvv")
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
