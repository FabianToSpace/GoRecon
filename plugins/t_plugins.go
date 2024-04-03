package plugins

import (
	"gorecon/config"
	"gorecon/logger"
)

var (
	Config config.Config
	Logger logger.ILogger
)

func Init() (config.Config, error) {
	if Config == (config.Config{}) {
		Config, err := config.GetConfig()
		if err != nil {
			return Config, err
		}
		Logger = logger.Logger(&Config)
		return Config, nil
	}
	return Config, nil
}
