package plugins

import (
	"bufio"
	"context"
	"gorecon/config"
	"gorecon/logger"
	"io"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func NmapTcpAll() PortScan {
	moduleName := "nmap-tcp-all"
	return PortScan{
		Name:        moduleName,
		Description: "nmap TCP all",
		Type:        "TCP",
		Tags:        []string{"default", "default-portscan"},
		Run: func(target string) []Service {
			services := make([]Service, 0)
			reader, writer := io.Pipe()

			cmdCtx, cmdDone := context.WithCancel(context.Background())

			scannerStopped := make(chan struct{})
			go func() {
				defer close(scannerStopped)

				scanner := bufio.NewScanner(reader)
				for scanner.Scan() {
					re := regexp.MustCompile(`^(\d+)\/((tcp|udp))(.*)open(\s*)([\w\-\/]+)(\s*)(.*)$`)
					matches := re.FindStringSubmatch(scanner.Text())
					if len(matches) > 0 {
						portString := matches[1]
						serviceName := matches[6]
						verion := matches[8]
						secure := strings.Contains(serviceName, "ssl") || strings.Contains(serviceName, "tls")

						if strings.HasPrefix(serviceName, "ssl/") || strings.HasPrefix(serviceName, "tls/") {
							serviceName = serviceName[4:]
						}

						logger.Logger().Info(moduleName, "Found open port: "+portString)

						// convert portString to int
						port, _ := strconv.Atoi(portString)
						services = append(services, Service{
							Target:   target,
							Protocol: "tcp",
							Port:     int(port),
							Name:     serviceName,
							Secure:   secure,
							Version:  verion,
						})
					}

				}
			}()

			cmd := exec.Command("nmap", target, "-sC", "-sV", "-p"+config.GetConfig().PortRange)
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
