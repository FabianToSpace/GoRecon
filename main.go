package main

import (
	"gorecon/config"
	"gorecon/logger"
	"gorecon/plugins"
	"os"
	"strconv"
	"time"
)

var (
	Target = ""
)

func main() {
	Target = os.Args[1]

	threads := config.GetConfig().Threads

	scanners := []plugins.PortScan{
		plugins.NmapTcpTop(),
		plugins.NmapUdpTop(),
		plugins.NmapTcpAll(),
	}
	serviceResults := make([]plugins.Service, 0)

	scanThreads := min(threads, len(scanners))
	logger.RunningTasks = len(scanners)

	sem := make(chan struct{}, scanThreads)

	ticker := make(chan struct{})
	go StartTicker(ticker)

	for _, scanner := range scanners {
		sem <- struct{}{}
		go func(s plugins.PortScan) {
			defer func() { <-sem }()
			done := make(chan []plugins.Service)
			go func() {
				logger.Logger().Start(s.Name, Target, "Starting "+s.Name)
				logger.ActiveTasks[s.Name] = true

				services := s.Run(Target)
				done <- services
			}()
			services := <-done
			logger.Logger().Done(s.Name, Target, "Done, found "+strconv.Itoa(len(services))+" services")

			logger.RunningTasks -= 1
			delete(logger.ActiveTasks, s.Name)

			serviceResults = append(serviceResults, services...)
		}(scanner)
	}

	for i := 0; i < scanThreads; i++ {
		sem <- struct{}{}
	}

	close(sem)
	close(ticker)

	logger.Logger().Done("Portscan", Target, "Found "+strconv.Itoa(len(serviceResults))+" services")
}

func StartTicker(quit chan struct{}) {
	ticker := time.NewTicker(10 * time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				logger.Logger().Ticker(Target)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}
