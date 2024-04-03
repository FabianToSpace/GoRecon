package plugins

import (
	"os"
	"testing"

	"github.com/FabianToSpace/GoRecon/config"
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
	if !cfg.Initialized && err == nil {
		t.Error("Config was not initialized as expected")
	}
	err = nil

	cfg, err = Init()
	if cfg.PortRange != "1-15" && err == nil {
		t.Error("Config was not initialized with the expected values")
	}
	err = nil

	Config = config.Config{PortRange: "1-65535", Initialized: true}
	cfg, err = Init()
	if cfg.PortRange != "1-65535" && err == nil {
		t.Error("Config was not initialized with the expected values")
	}
	err = nil
	Config = config.Config{}

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

func TestLoadPortScanners(t *testing.T) {
	// Test case where scanners match configured service scans
	Config.Initialized = true
	Config.Plugins.PortScans = []string{"nmap-tcp-all", "nmap-tcp-top"}
	expectedScanners := []PortScan{NmapTcpAll(), NmapTcpTop()}

	resultScanners := LoadPortScanners()

	if len(expectedScanners) != len(resultScanners) {
		t.Errorf("Expected %d scanners, but got %d", len(expectedScanners), len(resultScanners))
	}

	match := make(map[string]bool)

	for _, expected := range expectedScanners {
		for _, actual := range resultScanners {
			if expected.Name == actual.Name {
				match[expected.Name] = true
			}
		}
	}

	if len(match) != len(expectedScanners) {
		t.Errorf("Expected %d scanners, but got %d", len(expectedScanners), len(match))
	}

	// Test case where no scanners match configured service scans
	Config.Plugins.PortScans = []string{"NonExistentScanner"}
	expectedNoMatch := 0

	resultNoMatch := LoadPortScanners()

	if expectedNoMatch != len(resultNoMatch) {
		t.Errorf("Expected 0 scanners, but got %d", len(resultNoMatch))
	}
}

func TestLoadServiceScanners(t *testing.T) {
	// Test case where scanners match configured service scans
	Config.Initialized = true
	Config.Plugins.ServiceScans = []string{"dirbuster", "nmap-ftp"}
	expectedScanners := []ServiceScan{Dirbuster(), NmapFtp()}

	resultScanners := LoadServiceScanners()

	if len(expectedScanners) != len(resultScanners) {
		t.Errorf("Expected %d scanners, but got %d", len(expectedScanners), len(resultScanners))
	}

	match := make(map[string]bool)

	for _, expected := range expectedScanners {
		for _, actual := range resultScanners {
			if expected.Name == actual.Name {
				match[expected.Name] = true
			}
		}
	}

	if len(match) != len(expectedScanners) {
		t.Errorf("Expected %d scanners, but got %d", len(expectedScanners), len(match))
	}

	// Test case where no scanners match configured service scans
	Config.Plugins.ServiceScans = []string{"NonExistentScanner"}
	expectedNoMatch := 0

	resultNoMatch := LoadServiceScanners()

	if expectedNoMatch != len(resultNoMatch) {
		t.Errorf("Expected 0 scanners, but got %d", len(resultNoMatch))
	}
}
