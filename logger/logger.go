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
	Done  func(module, target, message string)
	Start func(module, target, message string)
}

func EnsureLength(name string, length int) string {
	for len(name) < length {
		name = name + " "
	}
	return name
}

func Printer(symbol, logtype, module, target, message string, colAtr color.Attribute, condition bool) {
	lock.Lock()
	defer lock.Unlock()
	module = EnsureLength(module, 20)
	logtype = EnsureLength(logtype, 5)
	if condition {
		c := color.New(colAtr, color.Bold)
		c.Printf("[%s] %s | ", symbol, logtype)
		c = color.New(color.FgMagenta, color.Bold)
		c.Printf("%s | ", target)
		c = color.New(color.FgGreen, color.Bold)
		c.Printf("%s\t", module)
		c.DisableColor()
		c.Printf("%s\n", message)
	}
}

func Logger() ILogger {
	return ILogger{
		Debug: func(module, target, message string) {
			Printer("=", "DEBUG", module, target, message, color.FgCyan, config.GetConfig().Debug)
		},
		Info: func(module, target, message string) {
			Printer("+", "INFO", module, target, message, color.FgCyan, true)
		},
		Warn: func(module, target, message string) {
			Printer("!", "WARN", module, target, message, color.FgYellow, true)
		},
		Error: func(module, target, message string) {
			Printer("X", "ERROR", module, target, message, color.FgRed, true)
		},
		Done: func(module, target, message string) {
			Printer("+", "DONE", module, target, message, color.FgGreen, true)
		},
		Start: func(module, target, message string) {
			Printer(">", "START", module, target, message, color.FgHiGreen, true)
		},
	}
}
