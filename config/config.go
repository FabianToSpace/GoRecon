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

	// check if file exists
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

	ctx := context.Background()
	if err := envconfig.ProcessWith(ctx, &envconfig.Config{
		Target:           &c,
		DefaultOverwrite: true,
	}); err != nil {
		return c, err
	}

	return c, nil
}
