package config

import (
	"context"
	"log"
	"os"

	envconfig "github.com/sethvargo/go-envconfig"
	"gopkg.in/yaml.v2"
)

type Config struct {
	PortRange    string `env:"PORT_RANGE"`
	OutputFormat string `env:"OUTPUT_FORMAT"`
	OutputFile   string `env:"OUTPUT_FILE"`
	Debug        bool   `env:"DEBUG"`
}

func GetConfig() Config {
	var c Config
	f, err := os.Open("config.yaml")

	if err != nil {
		log.Fatal(err)
	}

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&c)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	ctx := context.Background()
	if err := envconfig.ProcessWith(ctx, &envconfig.Config{
		Target:           &c,
		DefaultOverwrite: true,
	}); err != nil {
		log.Fatal(err)
	}

	return c
}
