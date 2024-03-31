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

func TestTokenizeArguments(t *testing.T) {
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
			expected: []string{"example.com", "token1", "token2", "token3", "/tmp/output_example.com.txt"},
			target:   "example.com",
		},
	}

	curDir, _ := os.Getwd()
	os.Chdir("/tmp")

	for _, tc := range testCases {
		result := tc.p.TokenizeArguments(tc.target)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v, but got %v", tc.expected, result)
		}
	}

	os.Chdir(curDir)
}
