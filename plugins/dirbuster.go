package plugins

import (
	"gorecon/logger"
)

func Dirbuster() ServiceScan {
	moduleName := "dirbuster"
	return ServiceScan{
		Name:        moduleName,
		Description: "Directory Buster",
		Tags:        []string{"default", "http"},
		Run: func(target string) []Service {
			logger.Logger().Start(moduleName, target, "Starting")

			logger.Logger().Done(moduleName, target, "Done")
			return []Service{}
		},
	}
}
