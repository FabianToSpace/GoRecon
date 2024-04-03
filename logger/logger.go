package logger

import (
	"fmt"
	"gorecon/config"
	"strings"
	"sync"
	"time"

	"github.com/mitchellh/colorstring"
)

var (
	lock   sync.Mutex
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

type ILogger struct {
	Config config.Config
}

func EnsureLength(name string, length int) string {
	for len(name) < length {
		name = name + " "
	}
	return name
}

func TimeNow(now func() time.Time) time.Time {
	if now == nil {
		now = time.Now
	}

	return now()
}

func (l ILogger) Printer(symbol, logtype, module, target, message string, color string, condition bool) {
	lock.Lock()
	defer lock.Unlock()
	module = EnsureLength(module, 35)
	logtype = EnsureLength(logtype, 5)
	target = EnsureLength(target, 22)

	if condition {
		fmt.Printf(colors.Color("["+color+"][%s] %s | [magenta]%s | [green]%s[reset]\t%s\n"), symbol, logtype, target, module, message)
	}
}

func (l ILogger) Debug(module, target, message string) {
	l.Printer("", "DEBUG", module, target, message, "cyan", l.Config.Debug)
}

func (l ILogger) Info(module, target, message string) {
	l.Printer("*", "INFO", module, target, message, "cyan", true)
}

func (l ILogger) Warn(module, target, message string) {
	l.Printer("!", "WARN", module, target, message, "yellow", true)
}

func (l ILogger) Error(module, target, message string) {
	l.Printer("X", "ERROR", module, target, message, "red", true)
}

func (l ILogger) Done(module, target, message string) {
	l.Printer("+", "DONE", module, target, message, "green", true)
}

func (l ILogger) Start(module, target, message string) {
	l.Printer(">", "START", module, target, message, "green", true)
}

func (l ILogger) Ticker(target string, t func() time.Time) {
	running := make([]string, 0)
	for k := range ActiveTasks {
		running = append(running, k)
	}

	if t == nil {
		t = time.Now
	}

	now := TimeNow(t).Format("15:04:05")
	fmt.Printf(
		colors.Color("[green][*] %s | [magenta]%s | [reset]Still running [yellow]%d[reset] Tasks\n[yellow]%s[reset]\n"),
		now, target, RunningTasks, strings.Join(running, ", "))
}

func Logger(cfg *config.Config) ILogger {
	if cfg == nil {
		cfg = &config.Config{}
	}

	return ILogger{
		Config: *cfg,
	}
}
