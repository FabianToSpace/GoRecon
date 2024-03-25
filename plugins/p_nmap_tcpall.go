package plugins

import (
	"gorecon/config"
)

func NmapTcpAll() PortScan {
	moduleName := "nmap-tcp-all"
	return PortScan{
		Name:         moduleName,
		Description:  "Nmap TCP all scan",
		Type:         "tcp",
		Tags:         []string{"default", "default-portscan"},
		Command:      "nmap",
		Arguments:    []string{"-sC", "-sV", "-p" + config.GetConfig().PortRange, "-vvvv"},
		TargetAppend: false,
	}
}
