package plugins

import (
	"os"
	"reflect"
	"testing"

	"github.com/FabianToSpace/GoRecon/config"
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
	tmpDir, _ := os.MkdirTemp("", "test")

	p := PortScan{}
	p.OutputFormat = "output_{{.Target}}.txt"
	os.Chdir(tmpDir)

	// Test when target is "example.com"
	target := "example.com"
	expected := tmpDir + "/output_example.com.txt"
	result := p.TokenizeOutput(target)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}

	// Test when target is "localhost"
	target = "localhost"
	expected = tmpDir + "/output_localhost.txt"
	result = p.TokenizeOutput(target)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}

	os.Chdir(curDir)
	os.RemoveAll(tmpDir)
}

func TestTokenizeArguments(t *testing.T) {
	curDir, _ := os.Getwd()
	tmpDir, _ := os.MkdirTemp("", "test")
	os.Chdir(tmpDir)

	Config = config.Config{PortRange: "1-65535"}
	testCases := []struct {
		p        PortScan
		expected []string
		target   string
	}{
		{
			p: PortScan{
				Arguments:    []string{"token1", "token2", "token3"},
				TargetAppend: false,
			},
			expected: []string{"example.com", "token1", "token2", "token3"},
			target:   "example.com",
		},
		{
			p: PortScan{
				Arguments:    []string{"token1", "token2", "token3"},
				TargetAppend: true,
			},
			expected: []string{"token1", "token2", "token3", "example.com"},
			target:   "example.com",
		},
		{
			p: PortScan{
				Arguments:        []string{"token1", "token2", "token3", "{{.PortRange}}"},
				ArgumentsInPlace: true,
			},
			expected: []string{"example.com", "token1", "token2", "token3", "1-65535"},
			target:   "example.com",
		},
		{
			p: PortScan{
				Arguments:    []string{"token1", "token2", "token3", "{{.OutputFile}}"},
				OutputFormat: "output_{{.Target}}.txt",
			},
			expected: []string{"example.com", "token1", "token2", "token3", tmpDir + "/output_example.com.txt"},
			target:   "example.com",
		},
	}

	for _, tc := range testCases {
		result := tc.p.TokenizeArguments(tc.target)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v, but got %v", tc.expected, result)
		}
	}

	os.Chdir(curDir)
	os.RemoveAll(tmpDir)
}
