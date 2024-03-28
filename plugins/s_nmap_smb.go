//nmap -sV -p21 -script="banner,(ftp* or ssl*) and not (brute or backdoor or libopie)" 21 10.10.127.244

package plugins

func NmapSmb() ServiceScan {
	moduleName := "nmap-smb"
	return ServiceScan{
		Name:             moduleName,
		Description:      "Nmap SMB Scripts",
		Tags:             []string{"default", "smb"},
		Command:          "nmap",
		Arguments:        []string{"-sV", "-p{{.Port}}", "-oN", "{{.OutputFile}}", "--script", "smb-enum-* and safe"},
		ArgumentsInPlace: true,
		TargetAppend:     true,
		MatchPattern:     "^(netbios-ssn|microsoft-ds|ldap|smb)",
		OutputFormat:     "results/{{.Target}}/scans/{{.Port}}-{{.Protocol}}/" + moduleName + ".txt",
		TargetFormat:     "{{.Target}}",
	}
}
