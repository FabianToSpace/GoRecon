package plugins

func NmapTcpTop() PortScan {
	moduleName := "nmap-tcp-top"
	return PortScan{
		Name:         moduleName,
		Description:  "Nmap TCP top scan",
		Type:         "tcp",
		Tags:         []string{"default", "default-portscan"},
		Command:      "nmap",
		Arguments:    []string{"-sC", "-sV", "-vvvv", "-oN", "{{.OutputFile}}"},
		OutputFormat: "results/{{.Target}}/scans/" + moduleName + ".txt",
	}
}
