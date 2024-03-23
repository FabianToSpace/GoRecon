package main

import (
	"gorecon/config"
	"gorecon/logger"
	"gorecon/plugins"
	"os"
	"strconv"
)

func main() {
	ip := os.Args[1]

	threads := config.GetConfig().Threads

	scanners := []plugins.PortScan{
		plugins.NmapTcpTop(),
		plugins.NmapUdpTop(),
		plugins.NmapTcpAll(),
	}
	serviceResults := make([]plugins.Service, 0)
	scanThreads := min(threads, len(scanners))

	sem := make(chan struct{}, scanThreads)
	for _, scanner := range scanners {
		sem <- struct{}{}
		go func(s plugins.PortScan) {
			defer func() { <-sem }()
			done := make(chan []plugins.Service)
			go func() {
				logger.Logger().Start(s.Name, ip, "Starting "+s.Name)
				services := s.Run(ip)
				done <- services
			}()
			services := <-done
			logger.Logger().Done(s.Name, ip, "Done, found "+strconv.Itoa(len(services))+" services")
			serviceResults = append(serviceResults, services...)
		}(scanner)
	}

	for i := 0; i < scanThreads; i++ {
		sem <- struct{}{}
	}

	close(sem)

	logger.Logger().Done("Portscan", ip, "Found "+strconv.Itoa(len(serviceResults))+" services")
}
