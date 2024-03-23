package main

import (
	"fmt"
	"gorecon/config"
	"gorecon/plugins"
	"os"
)

func main() {
	ip := os.Args[1]

	threads := config.GetConfig().Threads

	scanners := []plugins.PortScan{
		plugins.NmapTcpTop(),
		plugins.NmapUdpTop(),
		plugins.NmapTcpAll(),
	}

	scanThreads := min(threads, len(scanners))

	sem := make(chan struct{}, scanThreads)
	for _, scanner := range scanners {
		sem <- struct{}{}
		go func(s plugins.PortScan) {
			defer func() { <-sem }()
			services := s.Run(ip)
			fmt.Println(services)
		}(scanner)
	}

	for i := 0; i < scanThreads; i++ {
		sem <- struct{}{}
	}

	close(sem)
}
