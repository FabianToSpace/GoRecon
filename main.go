package main

import (
	"fmt"
	"gorecon/plugins"
	"os"
)

// Take an IP address as an positional parameter and run nmap against it
// further take
func main() {
	ip := os.Args[1]
	services := plugins.NmapTcpAll().Run(ip)
	fmt.Println(services)
}
