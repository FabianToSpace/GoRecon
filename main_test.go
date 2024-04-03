package main

import (
	"os"
	"testing"

	"github.com/FabianToSpace/GoRecon/config"
	"github.com/FabianToSpace/GoRecon/logger"
)

func TestMain(m *testing.M) {
	Config = config.Config{
		PortRange:    "1-65535",
		OutputFormat: "default",
		OutputFile:   "default",
		Debug:        false,
		Threads:      10,
	}

	Logger = logger.ILogger{
		Config: Config,
	}

	os.Exit(m.Run())
}

func TestCreatePaths(t *testing.T) {
	// Set up test environment
	oldDir, _ := os.Getwd()
	testDir, _ := os.MkdirTemp("", "test")
	defer os.RemoveAll(testDir)
	os.Chdir(testDir)

	// Call the function
	if err := CreatePaths(); err != nil {
		t.Errorf("Failed to create directories: %v", err)
	}

	// Verify that the directories were created
	_, err := os.Stat("results/" + Target + "/scans")
	if err != nil {
		t.Errorf("Failed to create directories: %v", err)
	}

	// Clean up test environment
	os.Chdir(oldDir)
}

func TestCreatePathsError(t *testing.T) {
	// Set up test environment
	oldDir, _ := os.Getwd()
	testDir, _ := os.MkdirTemp("", "test")

	os.Chdir(testDir)
	os.Mkdir("results", os.ModePerm)

	// Make the directory non-writable
	err := os.Chmod("results", 0o444)
	if err != nil {
		t.Fatalf("Failed to set directory permissions: %v", err)
	}

	// Call the function
	err = CreatePaths()

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	// Clean up test environment
	os.Chdir(oldDir)
	os.RemoveAll(testDir)
}
