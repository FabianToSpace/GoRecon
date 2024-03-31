package plugins

import (
	"reflect"
	"testing"
)

func TestNmapTcpTop(t *testing.T) {
	// Testing Name field
	ps := NmapTcpTop()
	expectedName := "nmap-tcp-top"
	if ps.Name != expectedName {
		t.Errorf("Expected Name to be %s, got %s", expectedName, ps.Name)
	}

	// Testing Description field
	expectedDescription := "Nmap TCP top scan"
	if ps.Description != expectedDescription {
		t.Errorf("Expected Description to be %s, got %s", expectedDescription, ps.Description)
	}

	// Testing Type field
	expectedType := "tcp"
	if ps.Type != expectedType {
		t.Errorf("Expected Type to be %s, got %s", expectedType, ps.Type)
	}

	// Testing Tags field
	expectedTags := []string{"default", "default-portscan"}
	if !reflect.DeepEqual(ps.Tags, expectedTags) {
		t.Errorf("Expected Tags to be %v, got %v", expectedTags, ps.Tags)
	}

	// Testing Command field
	expectedCommand := "nmap"
	if ps.Command != expectedCommand {
		t.Errorf("Expected Command to be %s, got %s", expectedCommand, ps.Command)
	}

	// Testing Arguments field
	expectedArguments := []string{"-sC", "-sV", "-vvvv", "-oN", "{{.OutputFile}}"}
	if !reflect.DeepEqual(ps.Arguments, expectedArguments) {
		t.Errorf("Expected Arguments to be %v, got %v", expectedArguments, ps.Arguments)
	}

	// Testing OutputFormat field
	expectedOutputFormat := "results/{{.Target}}/scans/nmap-tcp-top.txt"
	if ps.OutputFormat != expectedOutputFormat {
		t.Errorf("Expected OutputFormat to be %s, got %s", expectedOutputFormat, ps.OutputFormat)
	}
}
