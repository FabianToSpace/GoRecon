//nmap -sV -p21 -script="banner,(ftp* or ssl*) and not (brute or backdoor or libopie)" 21 10.10.127.244

package plugins

func Enum4Linux() ServiceScan {
	moduleName := "enum4linux"
	return ServiceScan{
		Name:             moduleName,
		Description:      "Enum4Linux Samba Enumeration",
		Tags:             []string{"default", "smb"},
		Command:          "enum4linux",
		Arguments:        []string{"-a"},
		ArgumentsInPlace: true,
		TargetAppend:     true,
		MatchPattern:     "^(netbios-ssn|microsoft-ds|ldap|smb)",
		OutputFormat:     "results/{{.Target}}/scans/{{.Port}}-{{.Protocol}}/" + moduleName + ".txt",
		TargetFormat:     "{{.Target}}",
		OutFile:          true,
	}
}
