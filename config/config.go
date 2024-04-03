package config

import (
	"context"
	"os"

	envconfig "github.com/sethvargo/go-envconfig"
	"gopkg.in/yaml.v2"
)

type Config struct {
	PortRange    string `env:"PORT_RANGE"`
	OutputFormat string `env:"OUTPUT_FORMAT"`
	OutputFile   string `env:"OUTPUT_FILE"`
	Debug        bool   `env:"DEBUG"`
	Threads      int    `env:"THREADS"`
	Plugins      Plugins
	Initialized  bool
}

type Plugins struct {
	PortScans    []string `env:"PORT_SCANS, delimiter=;"`
	ServiceScans []string `env:"SERVICE_SCANS, delimiter=;"`
}

var (
	AllowedCommands = []string{
		"nmap", "dirb", "feroxbuster", "whatweb", "nikto", "enum4linux",
	}
)

func GetConfig() (Config, error) {
	var c Config

	f, err := os.Open("config.yaml")
	if err != nil {
		return c, err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&c); err != nil {
		return c, err
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
