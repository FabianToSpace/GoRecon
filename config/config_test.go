package config

import (
	"os"
	"testing"
)

func TestGetConfig(t *testing.T) {
	// Test for successful reading of config file
	t.Run("Successful reading of config file", func(t *testing.T) {
		curdir, _ := os.Getwd()
		tmp, _ := os.MkdirTemp("", "configtest")
		os.Chdir(tmp)

		os.WriteFile("config.yaml", []byte(`portrange: 1
outputformat: test
outputfile: test.json
debug: false
threads: 1`), 0644)

		Config, err := GetConfig()
		if err != nil {
			t.Errorf("Error reading config file: %v", err)
		}
		if Config.PortRange != "1" ||
			Config.OutputFormat != "test" ||
			Config.OutputFile != "test.json" ||
			Config.Debug != false ||
			Config.Threads != 1 {
			t.Errorf("Invalid config values")
		}

		os.Chdir(curdir)
		os.RemoveAll(tmp)
	})

	// Test for handling error when opening config file
	t.Run("Error opening config file", func(t *testing.T) {
		curdir, _ := os.Getwd()
		tmp, _ := os.MkdirTemp("", "configtest")
		os.Chdir(tmp)

		_, err := GetConfig()
		if err == nil {
			t.Errorf("Expected error when opening config file: %v", err)
		}

		os.Chdir(curdir)
		os.RemoveAll(tmp)
	})

	// Test for handling error when decoding config file
	t.Run("Error decoding config file", func(t *testing.T) {
		curdir, _ := os.Getwd()
		tmp, _ := os.MkdirTemp("", "configtest")
		os.Chdir(tmp)

		os.WriteFile("config.yaml", []byte(`portrange: 1
	outputformat: test
		outputfile: test.json
debug: false
	threads: 1`), 0644)

		_, err := GetConfig()
		if err == nil {
			t.Errorf("Expected error when decoding config file: %v", err)
		}

		os.Chdir(curdir)
		os.RemoveAll(tmp)
	})

	// Test for handling error when processing config with environment variables
	t.Run("Error processing config with environment variables", func(t *testing.T) {
		curdir, _ := os.Getwd()
		tmp, _ := os.MkdirTemp("", "configtest")
		os.Chdir(tmp)

		os.WriteFile("config.yaml", []byte(`portrange: 1
outputformat: test
outputfile: test.json
debug: false
threads: 1`), 0644)

		os.Setenv("PORT_RANGE", "5")
		os.Setenv("OUTPUT_FORMAT", "env")
		os.Setenv("OUTPUT_FILE", "env.json")
		os.Setenv("DEBUG", "true")
		os.Setenv("THREADS", "5")

		Config, err := GetConfig()
		if err != nil {
			t.Errorf("Error reading config file: %v", err)
		}

		if Config.PortRange != "5" ||
			Config.OutputFormat != "env" ||
			Config.OutputFile != "env.json" ||
			Config.Debug != true ||
			Config.Threads != 5 {
			t.Errorf("Invalid config values")
		}

		os.Chdir(curdir)
		os.RemoveAll(tmp)
	})
}
