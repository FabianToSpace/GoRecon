package plugins

func NmapUdpTop() PortScan {
	moduleName := "nmap-udp-top"
	return PortScan{
		Name:         moduleName,
		Description:  "Nmap UDP top scan",
		Type:         "udp",
		Tags:         []string{"default", "default-portscan"},
		Command:      "nmap",
		Arguments:    []string{"-sU", "-A", "--top-ports", "100", "-oN", "{{.OutputFile}}"},
		TargetAppend: true,
		OutputFormat: "results/{{.Target}}/Scans/" + moduleName + ".txt",
	}
}
