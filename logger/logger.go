package logger

import (
	"gorecon/config"
	"sync"

	"github.com/fatih/color"
)

var lock sync.Mutex

type ILogger struct {
	Debug func(module, target, message string)
	Info  func(module, target, message string)
	Warn  func(module, target, message string)
	Error func(module, target, message string)
}

func Logger() ILogger {
	return ILogger{
		Debug: func(module, target, message string) {
			lock.Lock()
			defer lock.Unlock()
			if config.GetConfig().Debug {
				c := color.New(color.FgCyan, color.Bold)
				c.Print("[=] DEBUG:")
				c = color.New(color.FgMagenta, color.Bold)
				c.Printf("%s:", target)
				c = color.New(color.FgGreen, color.Bold)
				c.Printf("%s\t", module)
				c.DisableColor()
				c.Printf("%s\n", message)
			}
		},
		Info: func(module, target, message string) {
			lock.Lock()
			defer lock.Unlock()
			c := color.New(color.FgCyan, color.Bold)
			c.Print("[+] INFO:")
			c = color.New(color.FgMagenta, color.Bold)
			c.Printf("%s:", target)
			c = color.New(color.FgGreen, color.Bold)
			c.Printf("%s\t", module)
			c.DisableColor()
			c.Printf("%s\n", message)
		},
		Warn: func(module, target, message string) {
			lock.Lock()
			defer lock.Unlock()
			c := color.New(color.FgYellow, color.Bold)
			c.Print("[!] INFO|")
			c = color.New(color.FgMagenta, color.Bold)
			c.Printf("%s:", target)
			c = color.New(color.FgGreen, color.Bold)
			c.Printf("%s\t", module)
			c.DisableColor()
			c.Printf("%s\n", message)
		},
		Error: func(module, target, message string) {
			lock.Lock()
			defer lock.Unlock()
			c := color.New(color.FgRed, color.Bold)
			c.Print("[x] INFO:")
			c = color.New(color.FgMagenta, color.Bold)
			c.Printf("%s:", target)
			c = color.New(color.FgGreen, color.Bold)
			c.Printf("%s\t", module)
			c.DisableColor()
			c.Printf("%s\n", message)
		},
	}
}
