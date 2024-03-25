package plugins

import (
	"gorecon/logger"
	"regexp"
	"strconv"
	"strings"
)

type Service struct {
	Target   string // Target IP Address or Hostname
	Protocol string // TCP, UDP etc.
	Port     int    // Port Number where the service was found
	Name     string // Name of the service (e.g. http, ssh, ftp, etc.)
	Secure   bool   // Is using secure transportation
	Version  string // The version found (like nginx 1.2.3)
	Scheme   string // The Scheme (http, https, ftp, etc.)
}

func extractService(target, moduleName, line string) Service {
	re := regexp.MustCompile(`^(\d+)\/((tcp|udp))(.*)open(\s*)([\w\-\/]+)(\s*)(.*)$`)
	matches := re.FindStringSubmatch(line)
	if len(matches) > 0 {
		portString := matches[1]
		protocol := matches[2]
		serviceName := matches[6]
		verion := matches[8]
		secure := strings.Contains(serviceName, "ssl") || strings.Contains(serviceName, "tls") || strings.Contains(serviceName, "https")

		if strings.HasPrefix(serviceName, "ssl/") || strings.HasPrefix(serviceName, "tls/") {
			serviceName = serviceName[4:]
		}

		logger.Logger().Info(moduleName, target, "Found open port: "+portString)

		scheme := serviceName
		if secure {
			switch serviceName {
			case "http":
				scheme = "https"
			case "ftp":
				scheme = "sftp"
			}
		}

		port, _ := strconv.Atoi(portString)
		return Service{
			Target:   target,
			Protocol: protocol,
			Port:     int(port),
			Name:     serviceName,
			Secure:   secure,
			Version:  verion,
			Scheme:   scheme,
		}
	}

	return Service{}
}
