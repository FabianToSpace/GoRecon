package plugins

import (
	"gorecon/config"
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	Config = config.Config{}
	curDir, _ := os.Getwd()
	tmpDir, _ := os.MkdirTemp("", "test")
	os.Chdir(tmpDir)

	os.WriteFile("config.yaml", []byte(`portrange: 1-15
outputformat: test
outputfile: test.json
debug: false
threads: 1`), 0400)

	cfg, err := Init()
	if cfg == (config.Config{}) && err == nil {
		t.Error("Config was not initialized as expected")
	}
	err = nil

	cfg, err = Init()
	if cfg.PortRange != "1-15" && err == nil {
		t.Error("Config was not initialized with the expected values")
	}
	err = nil

	// Force Error
	os.Remove("config.yaml")
	os.WriteFile("config.yaml", []byte(`portrange: 1
	outputformat: test
		outputfile: test.json
debug: false
	threads: 1`), 0400)

	cfg, err = Init()

	if err == nil {
		t.Errorf("Expected error when opening config file: %v", err)
	}

	os.Chdir(curDir)
	os.RemoveAll(tmpDir)
}
