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
	Config       = config.Config{}
	Logger       = logger.ILogger{}
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
		plugins.NmapFtp(),
		plugins.Enum4Linux(),
		plugins.NmapSmb(),
	}
	Services = make([]plugins.Service, 0)
)

func main() {
	Target = os.Args[1]

	Config, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	Logger = logger.Logger(&Config)

	Threads = Config.Threads

	if err := CreatePaths(); err != nil {
		panic(err)
	}

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

	WriteServicesReport()

	StartServiceScanner()
}

func CreatePaths() error {
	curDir, _ := os.Getwd()
	if err := os.MkdirAll(curDir+"/results/"+Target+"/scans", os.ModePerm); err != nil {
		return err
	}

	return nil
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
		Logger.Start(scanner.Name, Target, "Starting "+scanner.Name)

		var mutex = &sync.Mutex{}
		mutex.Lock()
		logger.ActiveTasks[scanner.Name] = true
		mutex.Unlock()

		result := scanner.Run(Target)
		done <- result
	}()
	services := <-done
	Logger.Done(scanner.Name, Target, "Done, found "+strconv.Itoa(len(services))+" services")

	logger.RunningTasks -= 1
	delete(logger.ActiveTasks, scanner.Name)

	return services
}

func ServiceScannerRunner(scanner plugins.ServiceScan, service plugins.Service, wg *sync.WaitGroup) bool {
	defer wg.Done()
	taskname := fmt.Sprintf("%s-%s-%s-%d", scanner.Name, service.Name, service.Protocol, service.Port)
	done := make(chan bool)
	go func() {
		Logger.Start(taskname, Target, "Starting "+scanner.Name)

		var mutex = &sync.Mutex{}
		mutex.Lock()
		logger.ActiveTasks[taskname] = true
		mutex.Unlock()

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
				Logger.Ticker(Target, nil)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func WriteServicesReport() {
	filePath := "results/" + Target + "/services.txt"
	if outfile, err := os.Create(filePath); err != nil {
		panic(err)
	} else {
		defer outfile.Close()

		for _, service := range Services {
			serviceString := fmt.Sprintf(
				"Servicename:\t %s\nPort:\t\t\t %d\nProtocol:\t\t %s\nVersion:\t\t %s\nScheme:\t\t\t %s\nSecure:\t\t\t %t\n\n",
				service.Name, service.Port, service.Protocol, service.Version, service.Scheme, service.Secure)
			outfile.WriteString(serviceString)
		}
	}
}
