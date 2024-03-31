package plugins

import (
	"os"
	"reflect"
	"testing"
)

func TestReplaceInArguments(t *testing.T) {
	p := PortScan{Arguments: []string{"token1", "token2", "token3"}}
	token := "token"
	value := "replacement"

	result := p.ReplaceInArguments(token, value)

	expected := []string{"replacement1", "replacement2", "replacement3"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestTokenizeOutput(t *testing.T) {
	curDir, _ := os.Getwd()

	p := PortScan{}
	p.OutputFormat = "output_{{.Target}}.txt"
	os.Chdir("/tmp")

	// Test when target is "example.com"
	target := "example.com"
	expected := "/tmp/output_example.com.txt"
	result := p.TokenizeOutput(target)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}

	// Test when target is "localhost"
	target = "localhost"
	expected = "/tmp/output_localhost.txt"
	result = p.TokenizeOutput(target)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}

	os.Chdir(curDir)
}
