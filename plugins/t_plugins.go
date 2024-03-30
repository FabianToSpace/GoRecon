package plugins

import (
	"gorecon/config"
	"gorecon/logger"
)

var (
	Config config.Config
	Logger logger.ILogger
)

func Init() {
	if Config == (config.Config{}) {
		Config, err := config.GetConfig()
		if err != nil {
			panic(err)
		}
		Logger = logger.Logger(&Config)
	}
}
