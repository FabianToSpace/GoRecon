package messaging

import (
	"github.com/FabianToSpace/GoRecon/config"
	"github.com/FabianToSpace/GoRecon/logger"
)

var (
	Config config.Config
	Logger logger.ILogger
)

func Init() (config.Config, error) {
	if !Config.Initialized {
		Config, err := config.GetConfig()
		if err != nil {
			return Config, err
		}
		Logger = logger.Logger(&Config)
		Config.Initialized = true
	}
	return Config, nil
}
