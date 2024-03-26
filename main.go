package main

import (
	"fmt"
	"gorecon/config"
	"gorecon/logger"
	"gorecon/plugins"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	Target       = ""
	Threads      = 0
	PortScanners = []plugins.PortScan{
		plugins.NmapTcpTop(),
		plugins.NmapUdpTop(),
		plugins.NmapTcpAll(),
	}
	ServiceScanners = []plugins.ServiceScan{
		plugins.Dirbuster(),
		plugins.Whatweb(),
		plugins.Nikto(),
	}
	Services = make([]plugins.Service, 0)
)

func main() {
	Target = os.Args[1]
	Threads = config.GetConfig().Threads

	CreatePaths()

	StartPortScanner()

	uniqueServices := make(map[string]plugins.Service)
	for _, service := range Services {
		key := fmt.Sprintf("%s-%s-%d", service.Target, service.Protocol, service.Port)
		if _, ok := uniqueServices[key]; !ok {
			uniqueServices[key] = service
		}
	}
	Services = make([]plugins.Service, 0)
	for _, service := range uniqueServices {
		Services = append(Services, service)
	}

	StartServiceScanner()
}

func CreatePaths() {
	curDir, _ := os.Getwd()
	os.MkdirAll(curDir+"/results/"+Target+"/scans", os.ModePerm)
}

func StartPortScanner() {
	threads := min(Threads, len(PortScanners))
	logger.RunningTasks = len(PortScanners)

	sem := make(chan struct{}, threads)

	ticker := make(chan struct{})
	go StartTicker(ticker)

	for _, scanner := range PortScanners {
		sem <- struct{}{}
		go func() {
			result := PortScannerRunner(scanner)
			Services = append(Services, result...)
			<-sem
		}()
	}

	for i := 0; i < threads; i++ {
		sem <- struct{}{}
	}

	close(sem)
	ticker <- struct{}{}
	close(ticker)
}

func StartServiceScanner() {
	threads := min(Threads, len(ServiceScanners)*len(Services))
	logger.RunningTasks = len(ServiceScanners) * len(Services)

	sem := make(chan struct{}, threads)

	ticker := make(chan struct{})
	go StartTicker(ticker)

	var wg sync.WaitGroup

	for _, scanner := range ServiceScanners {
		for _, service := range Services {
			wg.Add(1)
			sem <- struct{}{}
			go func() {
				go ServiceScannerRunner(scanner, service, &wg)
				<-sem
			}()
		}
	}
	wg.Wait()
	for i := 0; i < threads; i++ {
		sem <- struct{}{}
	}

	close(sem)
	ticker <- struct{}{}
	close(ticker)
}

func PortScannerRunner(scanner plugins.PortScan) []plugins.Service {
	done := make(chan []plugins.Service)
	go func() {
		logger.Logger().Start(scanner.Name, Target, "Starting "+scanner.Name)
		logger.ActiveTasks[scanner.Name] = true

		result := scanner.Run(Target)
		done <- result
	}()
	services := <-done
	logger.Logger().Done(scanner.Name, Target, "Done, found "+strconv.Itoa(len(services))+" services")

	logger.RunningTasks -= 1
	delete(logger.ActiveTasks, scanner.Name)

	return services
}

func ServiceScannerRunner(scanner plugins.ServiceScan, service plugins.Service, wg *sync.WaitGroup) bool {
	defer wg.Done()
	taskname := fmt.Sprintf("%s-%s-%s-%d", scanner.Name, service.Name, service.Protocol, service.Port)
	done := make(chan bool)
	go func() {
		logger.Logger().Start(taskname, Target, "Starting "+scanner.Name)
		logger.ActiveTasks[taskname] = true

		result := scanner.Run(service)
		done <- result
	}()

	res := <-done
	logger.RunningTasks -= 1
	delete(logger.ActiveTasks, taskname)
	return res
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
