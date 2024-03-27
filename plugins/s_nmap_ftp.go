//nmap -sV -p21 -script="banner,(ftp* or ssl*) and not (brute or backdoor or libopie)" 21 10.10.127.244

package plugins

func NmapFtp() ServiceScan {
	moduleName := "nmap-ftp"
	return ServiceScan{
		Name:             moduleName,
		Description:      "Nmap FTP Scripts",
		Tags:             []string{"default", "ftp"},
		Command:          "nmap",
		Arguments:        []string{"-sV", "-p{{.Port}}", "-oN", "{{.OutputFile}}", "--script='banner,(ftp* or ssl*) and safe'"},
		ArgumentsInPlace: true,
		TargetAppend:     true,
		MatchPattern:     "^ftp",
		OutputFormat:     "results/{{.Target}}/scans/{{.Port}}-{{.Protocol}}/" + moduleName + ".txt",
		TargetFormat:     "{{.Target}}",
	}
}
