package plugins

func Whatweb() ServiceScan {
	moduleName := "whatweb"
	return ServiceScan{
		Name:         moduleName,
		Description:  "Directory Buster",
		Tags:         []string{"default", "http"},
		Command:      "whatweb",
		Arguments:    []string{"--color=never", "--no-errors", "-a 3", "-v"},
		TargetFormat: "{{.Scheme}}://{{.Target}}:{{.Port}}",
		TargetAppend: true,
		MatchPattern: "^http",
		OutputFormat: "results/{{.Target}}/scans/{{.Port}}-{{.Protocol}}/" + moduleName + ".txt",
		OutFile:      true,
	}
}
