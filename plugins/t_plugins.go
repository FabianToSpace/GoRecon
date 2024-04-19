package plugins

import (
	"strconv"

	"github.com/FabianToSpace/GoRecon/config"
	"github.com/FabianToSpace/GoRecon/logger"
	"github.com/FabianToSpace/GoRecon/messaging"
)

var (
	Config           config.Config
	Logger           logger.ILogger
	RabbitConnection = messaging.RabbitConnection{}
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
	RabbitConnection = messaging.RabbitConnection{
		ConnectionInfo: messaging.ConnectionInfo{
			User:     Config.Messaging.RabbitMq.User,
			Password: Config.Messaging.RabbitMq.Password,
			Host:     Config.Messaging.RabbitMq.Host,
			Port:     Config.Messaging.RabbitMq.Port,
		},
	}
	return Config, nil
}

func LoadPortScanners() []PortScan {
	Config, _ = Init()

	allScanners := []PortScan{
		NmapTcpTop(),
		NmapUdpTop(),
		NmapTcpAll(),
	}

	scanners := make([]PortScan, 0)

	for _, scanner := range allScanners {
		for _, configScanner := range Config.Plugins.PortScans {
			if scanner.Name == configScanner {
				scanners = append(scanners, scanner)
			}
		}
	}

	Logger.Debug("Plugins", "", "Loaded "+strconv.Itoa(len(scanners))+" Port scanners")

	return scanners
}

func LoadServiceScanners() []ServiceScan {
	Config, _ = Init()
	allScanners := []ServiceScan{
		Dirbuster(),
		Whatweb(),
		Nikto(),
		NmapFtp(),
		Enum4Linux(),
		NmapSmb(),
	}

	scanners := make([]ServiceScan, 0)

	for _, scanner := range allScanners {
		for _, configScanner := range Config.Plugins.ServiceScans {
			if scanner.Name == configScanner {
				scanners = append(scanners, scanner)
			}
		}
	}

	Logger.Debug("Plugins", "", "Loaded "+strconv.Itoa(len(scanners))+" Service scanners")

	return scanners
}
