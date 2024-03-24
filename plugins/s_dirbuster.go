package plugins

func Dirbuster() ServiceScan {
	moduleName := "dirbuster"
	return ServiceScan{
		Name:        moduleName,
		Description: "Directory Buster",
		Tags:        []string{"default", "http"},
		Run: func(service Service) bool {
			return false
		},
	}
}
