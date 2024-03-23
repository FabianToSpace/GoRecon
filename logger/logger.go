package logger

import (
	"gorecon/config"

	"github.com/fatih/color"
)

type ILogger struct {
	Debug func(module, message string)
	Info  func(module, message string)
	Warn  func(module, message string)
	Error func(module, message string)
}

func Logger() ILogger {
	return ILogger{
		Debug: func(module, message string) {
			if config.GetConfig().Debug {
				c := color.New(color.FgCyan)
				c.Print("[=] DEBUG: ")
				c.DisableColor()
				c.Printf("%s: %s\n", module, message)
			}
		},
		Info: func(module, message string) {
			c := color.New(color.FgMagenta)
			c.Print("[+] INFO: ")
			c.DisableColor()
			c.Printf("%s: %s\n", module, message)
		},
		Warn: func(module, message string) {
			c := color.New(color.FgYellow)
			c.Print("[!] INFO: ")
			c.DisableColor()
			c.Printf("%s: %s\n", module, message)
		},
		Error: func(module, message string) {
			c := color.New(color.FgRed)
			c.Print("[x] INFO: ")
			c.DisableColor()
			c.Printf("%s: %s\n", module, message)
		},
	}
}
