package config

import (
	"context"
	"os"

	envconfig "github.com/sethvargo/go-envconfig"
	"gopkg.in/yaml.v2"
)

type Config struct {
	PortRange    string `env:"PORT_RANGE, default=1-65535"`
	OutputFormat string `env:"OUTPUT_FORMAT, default=json"`
	OutputFile   string `env:"OUTPUT_FILE, default=recon.json"`
	Debug        bool   `env:"DEBUG, default=false"`
	Threads      int    `env:"THREADS, default=10"`
	Plugins      Plugins
	Messaging    Messaging
	Initialized  bool
}

type Plugins struct {
	PortScans    []string `env:"PORT_SCANS, delimiter=;"`
	ServiceScans []string `env:"SERVICE_SCANS, delimiter=;"`
}

type Messaging struct {
	RabbitMq RabbitMqConfig
}

type RabbitMqConfig struct {
	Host      string `env:"RABBITMQ_HOST, default=127.0.0.1"`
	Port      string `env:"RABBITMQ_PORT, default=5672"`
	User      string `env:"RABBITMQ_USER, default=guest"`
	Password  string `env:"RABBITMQ_PASSWORD, default=guest"`
	Exchange  string `env:"RABBITMQ_EXCHANGE, default=gorecon"`
	Sender    bool   `env:"RABBITMQ_SENDER, default=false"`
	Receiver  bool   `env:"RABBITMQ_RECEIVER, default=false"`
	QueueName string `env:"RABBITMQ_QUEUE_NAME, default=recon"`
}

var (
	AllowedCommands = []string{
		"nmap", "dirb", "feroxbuster", "whatweb", "nikto", "enum4linux-ng",
	}
)

func GetConfig() (Config, error) {
	var c Config

	skipConfigFile := false

	if os.Getenv("SKIP_CONFIG_FILE") == "true" {
		skipConfigFile = true
	}

	// check if file exists
	if !skipConfigFile {
		if _, err := os.Stat("config.yaml"); err == nil {
			f, err := os.Open("config.yaml")
			if err != nil {
				return c, err
			}
			defer f.Close()

			decoder := yaml.NewDecoder(f)
			if err := decoder.Decode(&c); err != nil {
				return c, err
			}
		}
	}

	ctx := context.Background()
	if err := envconfig.ProcessWith(ctx, &envconfig.Config{
		Target:           &c,
		DefaultOverwrite: true,
	}); err != nil {
		return c, err
	}

	return c, nil
}
