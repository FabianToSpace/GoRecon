package plugins

func Nikto() ServiceScan {
	moduleName := "nikto"
	return ServiceScan{
		Name:          moduleName,
		Description:   "Nikto Web Server Scanner",
		Tags:          []string{"default", "http"},
		Command:       "nikto",
		Arguments:     []string{"-host", "{{.TargetPos}}", "-ask=no", "-Tuning=x4567890ac", "-nointeractive", "-r", "-e", "-o", "{{.OutputFile}}"},
		TargetFormat:  "{{.Scheme}}://{{.Target}}:{{.Port}}",
		TargetInplace: true,
		MatchPattern:  "^http",
		OutputFormat:  "results/{{.Target}}/Scans/{{.Port}}-{{.Protocol}}/" + moduleName + ".txt",
		OutFile:       true,
	}
}
