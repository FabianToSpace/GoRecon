package plugins

import (
	"gorecon/logger"
	"regexp"
	"strconv"
	"strings"
)

type Service struct {
	Target   string
	Protocol string
	Port     int
	Name     string
	Secure   bool
	Version  string
}

func extractService(target, moduleName, line string) Service {
	re := regexp.MustCompile(`^(\d+)\/((tcp|udp))(.*)open(\s*)([\w\-\/]+)(\s*)(.*)$`)
	matches := re.FindStringSubmatch(line)
	if len(matches) > 0 {
		portString := matches[1]
		serviceName := matches[6]
		verion := matches[8]
		secure := strings.Contains(serviceName, "ssl") || strings.Contains(serviceName, "tls")

		if strings.HasPrefix(serviceName, "ssl/") || strings.HasPrefix(serviceName, "tls/") {
			serviceName = serviceName[4:]
		}

		logger.Logger().Info(moduleName, target, "Found open port: "+portString)

		port, _ := strconv.Atoi(portString)
		return Service{
			Target:   target,
			Protocol: "tcp",
			Port:     int(port),
			Name:     serviceName,
			Secure:   secure,
			Version:  verion,
		}
	}

	return Service{}
}
