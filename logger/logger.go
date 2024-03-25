package logger

import (
	"fmt"
	"gorecon/config"
	"strings"
	"sync"
	"time"

	"github.com/mitchellh/colorstring"
)

var lock sync.Mutex

type ILogger struct {
	Debug  func(module, target, message string)
	Info   func(module, target, message string)
	Warn   func(module, target, message string)
	Error  func(module, target, message string)
	Done   func(module, target, message string)
	Start  func(module, target, message string)
	Ticker func(target string)
}

func EnsureLength(name string, length int) string {
	for len(name) < length {
		name = name + " "
	}
	return name
}

var (
	colors = colorstring.Colorize{
		Colors: map[string]string{
			"black":   "1;30",
			"red":     "1;31",
			"green":   "1;32",
			"yellow":  "1;33",
			"blue":    "1;34",
			"magenta": "1;35",
			"cyan":    "1;36",
			"white":   "1;37",
			"reset":   "0",
		},
		Reset: true,
	}

	RunningTasks = 0
	ActiveTasks  = make(map[string]bool)
)

func Printer(symbol, logtype, module, target, message string, color string, condition bool) {
	lock.Lock()
	defer lock.Unlock()
	module = EnsureLength(module, 35)
	logtype = EnsureLength(logtype, 5)
	target = EnsureLength(target, 22)

	if condition {
		fmt.Printf(colors.Color("["+color+"][%s] %s | [magenta]%s | [green]%s[reset]\t%s\n"), symbol, logtype, target, module, message)
	}
}

func Logger() ILogger {
	return ILogger{
		Debug: func(module, target, message string) {
			Printer("", "DEBUG", module, target, message, "cyan", config.GetConfig().Debug)
		},
		Info: func(module, target, message string) {
			Printer("*", "INFO", module, target, message, "cyan", true)
		},
		Warn: func(module, target, message string) {
			Printer("!", "WARN", module, target, message, "yellow", true)
		},
		Error: func(module, target, message string) {
			Printer("X", "ERROR", module, target, message, "red", true)
		},
		Done: func(module, target, message string) {
			Printer("+", "DONE", module, target, message, "green", config.GetConfig().Debug)
		},
		Start: func(module, target, message string) {
			Printer(">", "START", module, target, message, "green", config.GetConfig().Debug)
		},
		Ticker: func(target string) {
			running := make([]string, 0)
			for k := range ActiveTasks {
				running = append(running, k)
			}

			now := time.Now().Format("15:04:05")
			fmt.Printf(
				colors.Color("[green][*] %s | [magenta]%s | [reset]Still running [yellow]%d[reset] Tasks\n[yellow]%s[reset]\n"),
				now, target, RunningTasks, strings.Join(running, ", "))
		},
	}
}
