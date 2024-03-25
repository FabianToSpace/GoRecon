package plugins

func Dirbuster() ServiceScan {
	moduleName := "dirbuster"
	return ServiceScan{
		Name:          moduleName,
		Description:   "Directory Buster",
		Tags:          []string{"default", "http"},
		Command:       "feroxbuster",
		Arguments:     []string{"-u {{.TargetPos}}", "-v", "-k", "-q", "-r", "-e"},
		TargetFormat:  "{{.Scheme}}://{{.Target}}:{{.Port}}",
		TargetInplace: true,
		MatchPattern:  "^http",
	}
}
