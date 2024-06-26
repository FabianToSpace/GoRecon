package config

import (
	"os"
	"testing"
)

func TestGetConfig(t *testing.T) {
	// Config not existing
	t.Run("Config not existing", func(t *testing.T) {
		curdir, _ := os.Getwd()
		tmp, _ := os.MkdirTemp("", "configtest")
		os.Chdir(tmp)

		_, err := GetConfig()
		if err != nil {
			t.Error("Expected no error when config file is not existing")
		}

		os.Chdir(curdir)
		os.RemoveAll(tmp)
	})

	// Test for successful reading of config file
	t.Run("Successful reading of config file", func(t *testing.T) {
		curdir, _ := os.Getwd()
		tmp, _ := os.MkdirTemp("", "configtest")
		os.Chdir(tmp)

		os.WriteFile("config.yaml", []byte(`portrange: 1
outputformat: test
outputfile: test.json
debug: false
threads: 1
plugins:
  portscans:
    - nmap-tcp-all
  servicescans:
    - whatweb`), 0400)

		Config, err := GetConfig()
		if err != nil {
			t.Errorf("Error reading config file: %v", err)
		}
		if Config.PortRange != "1" ||
			Config.OutputFormat != "test" ||
			Config.OutputFile != "test.json" ||
			Config.Debug != false ||
			Config.Threads != 1 ||
			len(Config.Plugins.PortScans) != 1 ||
			len(Config.Plugins.ServiceScans) != 1 {
			t.Errorf("Invalid config values")
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
	threads: 1`), 0400)

		_, err := GetConfig()
		if err == nil {
			t.Errorf("Expected error when decoding config file: %v", err)
		}

		os.Chdir(curdir)
		os.RemoveAll(tmp)
	})

	// Test for handling error when processing config with environment variables
	t.Run("Overwriting config with environment variables", func(t *testing.T) {
		curdir, _ := os.Getwd()
		tmp, _ := os.MkdirTemp("", "configtest")
		os.Chdir(tmp)

		os.WriteFile("config.yaml", []byte(`portrange: 1
outputformat: test
outputfile: test.json
debug: false
threads: 1`), 0400)

		os.Setenv("PORT_RANGE", "5")
		os.Setenv("OUTPUT_FORMAT", "env")
		os.Setenv("OUTPUT_FILE", "env.json")
		os.Setenv("DEBUG", "true")
		os.Setenv("THREADS", "5")
		os.Setenv("PORT_SCANS", "nmap-tcp-all;nmap-tcp-top")
		os.Setenv("SERVICE_SCANS", "whatweb;nikto")

		Config, err := GetConfig()
		if err != nil {
			t.Errorf("Error reading config file: %v", err)
		}

		if Config.PortRange != "5" ||
			Config.OutputFormat != "env" ||
			Config.OutputFile != "env.json" ||
			Config.Debug != true ||
			Config.Threads != 5 ||
			len(Config.Plugins.PortScans) != 2 ||
			len(Config.Plugins.ServiceScans) != 2 {
			t.Errorf("Invalid config values")
		}

		resetEnv()

		os.Chdir(curdir)
		os.RemoveAll(tmp)
	})

	// Test for handling error when processing config with environment variables
	t.Run("Skip Configfile Loading", func(t *testing.T) {
		curdir, _ := os.Getwd()
		tmp, _ := os.MkdirTemp("", "configtest")
		os.Chdir(tmp)

		os.WriteFile("config.yaml", []byte(`portrange: 1
outputformat: test
outputfile: test.json
debug: false
threads: 1`), 0400)

		if err := os.Setenv("SKIP_CONFIG_FILE", "true"); err != nil {
			t.Errorf("Error setting environment variable: %v", err)
		}

		Config, err := GetConfig()
		if err != nil {
			t.Errorf("Error reading config file: %v", err)
		}

		if Config.PortRange != "1-65535" ||
			Config.OutputFormat != "json" ||
			Config.OutputFile != "recon.json" ||
			Config.Debug != false ||
			Config.Threads != 10 ||
			len(Config.Plugins.PortScans) != 0 ||
			len(Config.Plugins.ServiceScans) != 0 {
			t.Errorf("Invalid config values")
			t.Errorf("Config: %v", Config)
		}

		resetEnv()

		os.Chdir(curdir)
		os.RemoveAll(tmp)
	})
}

func resetEnv() {
	os.Unsetenv("PORT_RANGE")
	os.Unsetenv("OUTPUT_FORMAT")
	os.Unsetenv("OUTPUT_FILE")
	os.Unsetenv("DEBUG")
	os.Unsetenv("THREADS")
	os.Unsetenv("PORT_SCANS")
	os.Unsetenv("SERVICE_SCANS")
	os.Unsetenv("SKIP_CONFIG_FILE")
}
