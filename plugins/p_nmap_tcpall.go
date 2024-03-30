package plugins

func NmapTcpAll() PortScan {
	moduleName := "nmap-tcp-all"
	return PortScan{
		Name:         moduleName,
		Description:  "Nmap TCP all scan",
		Type:         "tcp",
		Tags:         []string{"default", "default-portscan"},
		Command:      "nmap",
		Arguments:    []string{"-sC", "-sV", "-p{{.PortRange}}", "-vvvv", "-oN", "{{.OutputFile}}"},
		OutputFormat: "results/{{.Target}}/scans/" + moduleName + ".txt",
	}
}
