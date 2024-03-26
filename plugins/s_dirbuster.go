package plugins

func Dirbuster() ServiceScan {
	moduleName := "dirbuster"
	return ServiceScan{
		Name:          moduleName,
		Description:   "Directory Buster",
		Tags:          []string{"default", "http"},
		Command:       "feroxbuster",
		Arguments:     []string{"-u", "{{.TargetPos}}", "-v", "-k", "-q", "-r", "-e", "-o", "{{.OutputFile}}"},
		TargetFormat:  "{{.Scheme}}://{{.Target}}:{{.Port}}",
		TargetInplace: true,
		MatchPattern:  "^http",
		OutputFormat:  "results/{{.Target}}/Scans/{{.Port}}-{{.Protocol}}/" + moduleName + ".txt",
	}
}
