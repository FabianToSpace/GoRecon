package plugins

import (
	"os"
	"reflect"
	"testing"
)

func TestServiceScanTokenizeOutput(t *testing.T) {
	curDir, _ := os.Getwd()
	tmpDir, _ := os.MkdirTemp("", "test")
	os.Chdir(tmpDir)

	s := ServiceScan{}
	service := Service{Target: "example.com", Port: 80, Protocol: "HTTP"}

	expectedOutput := tmpDir + "/path/to/output/example.com_80_HTTP"

	s.OutputFormat = "path/to/output/{{.Target}}_{{.Port}}_{{.Protocol}}"
	result := s.TokenizeOutput(service)

	if result != expectedOutput {
		t.Errorf("Expected: %s, but got: %s", expectedOutput, result)
	}

	os.Chdir(curDir)
	os.RemoveAll(tmpDir)
}

func TestMatchCondition(t *testing.T) {
	tests := []struct {
		name     string
		pattern  string
		service  Service
		expected bool
	}{
		{
			name:     "Matching service name",
			pattern:  "service",
			service:  Service{Name: "myservice"},
			expected: true,
		},
		{
			name:     "Non-matching service name",
			pattern:  "other",
			service:  Service{Name: "myservice"},
			expected: false,
		},
		{
			name:     "Empty service name",
			pattern:  "service",
			service:  Service{Name: ""},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := ServiceScan{MatchPattern: tt.pattern}
			result := s.MatchCondition(tt.service)
			if result != tt.expected {
				t.Errorf("Expected %v, but got %v", tt.expected, result)
			}
		})
	}
}

func TestServiceScanReplaceInArguments(t *testing.T) {
	s := ServiceScan{
		Arguments: []string{"Hello, {name}!", "How are you, {name}?"},
	}

	// Test case 1: Replace {name} token with "John" in the first argument
	expected := []string{"Hello, John!", "How are you, John?"}
	result := s.ReplaceInArguments("{name}", "John")
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case 1 failed. Expected %v, but got %v", expected, result)
	}

	s = ServiceScan{
		Arguments: []string{"Hello, {name}!", "How are you, {name}?"},
	}
	// Test case 2: Replace {name} token with "Jane" in both arguments
	expected = []string{"Hello, Jane!", "How are you, Jane?"}
	result = s.ReplaceInArguments("{name}", "Jane")
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case 2 failed. Expected %v, but got %v", expected, result)
	}
}

func TestEnsurePath(t *testing.T) {
	curDir, _ := os.Getwd()
	tmpDir, _ := os.MkdirTemp("", "test")
	os.Chdir(tmpDir)

	s := ServiceScan{OutputFormat: "{{.Target}}/output.txt"}
	service := Service{Target: "example.com"}
	result, err := s.EnsurePath(service)
	if !result && err != nil {
		t.Error("Expected true, got false")
	}

	os.Chdir(curDir)
	os.RemoveAll(tmpDir)
}

func TestEnsurePathError(t *testing.T) {
	curDir, _ := os.Getwd()
	tmpDir, _ := os.MkdirTemp("", "test")
	os.Chdir("tmpDir")

	s := ServiceScan{OutputFormat: "{{.Target}}/output.txt"}

	// make tmpDir write protected
	os.Chmod(tmpDir, 0000)
	service := Service{Target: "example2.com"}
	result, err := s.EnsurePath(service)
	if result || err == nil {
		t.Error("Expected false, got true")
		t.Errorf("Error: %v", err)
		t.Errorf("Result: %v", result)
	}

	os.Chdir(curDir)
	os.RemoveAll(tmpDir)
}

func TestServiceScanTokenizeArguments(t *testing.T) {
	s := ServiceScan{
		TargetFormat:     "http://{{.Target}}:{{.Port}}",
		OutputFormat:     "output{{.Target}}.txt",
		Arguments:        []string{},
		ArgumentsInPlace: true,
		TargetInplace:    false,
		TargetAppend:     true,
	}

	service := Service{
		Target: "example.com",
		Port:   8080,
		Scheme: "https",
	}

	expected := []string{"http://example.com:8080"}
	result := s.TokenizeArguments(service)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
