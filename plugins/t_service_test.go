package plugins

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

func TestExtractService(t *testing.T) {
	tests := []struct {
		name     string
		target   string
		module   string
		line     string
		expected Service
	}{
		{
			name:     "Valid TCP service",
			target:   "127.0.0.1",
			module:   "TestModule",
			line:     "123/tcp    open  http    Apache httpd",
			expected: Service{Target: "127.0.0.1", Protocol: "tcp", Port: 123, Name: "http", Secure: false, Version: "Apache httpd", Scheme: "http"},
		},
		{
			name:     "Valid UDP service",
			target:   "localhost",
			module:   "TestModule",
			line:     "456/udp    open  domain  ISC BIND",
			expected: Service{Target: "localhost", Protocol: "udp", Port: 456, Name: "domain", Secure: false, Version: "ISC BIND", Scheme: "domain"},
		},
		{
			name:     "Valid FTP service",
			target:   "example.com",
			module:   "TestModule",
			line:     "789/tcp    open  ssl/ftp    vsftpd 2.3.4",
			expected: Service{Target: "example.com", Protocol: "tcp", Port: 789, Name: "ftp", Secure: true, Version: "vsftpd 2.3.4", Scheme: "sftp"},
		},
		{
			name:     "Secure service",
			target:   "example.com",
			module:   "TestModule",
			line:     "789/tcp    open  ssl/http    Apache httpd",
			expected: Service{Target: "example.com", Protocol: "tcp", Port: 789, Name: "http", Secure: true, Version: "Apache httpd", Scheme: "https"},
		},
		{
			name:     "Invalid line format",
			target:   "example.com",
			module:   "TestModule",
			line:     "invalid line format",
			expected: Service{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractService(tt.target, tt.module, tt.line)
			if result != tt.expected {
				t.Errorf("Expected %+v, but got %+v", tt.expected, result)
			}
		})
	}
}
