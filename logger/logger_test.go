package logger

import (
	"fmt"
	"io"
	"os"
	"testing"
	"time"

	"github.com/FabianToSpace/GoRecon/config"
)

func TestEnsureLength_ValidCases(t *testing.T) {
	testCases := []struct {
		name     string
		length   int
		expected string
	}{
		{"hello world", 10, "hello world"},
		{"hello", 10, "hello     "},
		{"", 5, "     "},
	}

	for _, tc := range testCases {
		result := EnsureLength(tc.name, tc.length)
		if result != tc.expected {
			t.Errorf("EnsureLength(%q, %d) = %q; expected %q", tc.name, tc.length, result, tc.expected)
		}
	}
}

func TestEnsureLength_NegativeLength(t *testing.T) {
	name := "hello"
	length := -2
	expected := "hello"
	result := EnsureLength(name, length)
	if result != expected {
		t.Errorf("EnsureLength(%q, %d) = %q; expected %q", name, length, result, expected)
	}
}

func TestEnsureLength_ZeroLength(t *testing.T) {
	name := "hello"
	length := 0
	expected := "hello"
	result := EnsureLength(name, length)
	if result != expected {
		t.Errorf("EnsureLength(%q, %d) = %q; expected %q", name, length, result, expected)
	}
}

func TestTimeNow(t *testing.T) {
	setTime := time.Unix(1234567890, 0)
	result := TimeNow(func() time.Time { return setTime })
	if result != setTime {
		t.Errorf("TimeNow() = %v; expected %v", result, setTime)
	}

	result = TimeNow(nil)
	if result.IsZero() {
		t.Errorf("TimeNow() = %v; expected non-zero", result)
	}
}

func TestPrinter_ConditionTrue(t *testing.T) {
	symbol := "+"
	logtype := "INFO"
	module := "TestModule"
	target := "TestTarget"
	message := "TestMessage"
	color := "green"
	condition := true

	out := capturePrintOutput(symbol, logtype, module, target, message, color, condition)

	module = EnsureLength(module, 35)
	logtype = EnsureLength(logtype, 5)
	target = EnsureLength(target, 22)

	expectedOutput := fmt.Sprintf(colors.Color("[green][%s] %s | [magenta]%s | [green]%s[reset]\t%s\n"), symbol, logtype, target, module, message)

	if string(out) != expectedOutput {
		t.Errorf("Unexpected output: got %q, want %q", string(out), expectedOutput)
	}
}

func TestPrinter_ConditionFalse(t *testing.T) {
	symbol := "+"
	logtype := "INFO"
	module := "TestModule"
	target := "TestTarget"
	message := "TestMessage"
	color := "green"
	condition := false

	out := capturePrintOutput(symbol, logtype, module, target, message, color, condition)

	if len(out) != 0 {
		t.Errorf("Unexpected output: got %q, want empty", string(out))
	}
}

func TestLoggerFuncs(t *testing.T) {
	logger := ILogger{Config: config.Config{Debug: true}}
	module := "TestModule"
	target := "TestTarget"
	message := "TestMessage"

	tests := []struct {
		symbol    string
		logtype   string
		color     string
		condition bool
	}{
		{
			symbol:    "",
			logtype:   "DEBUG",
			color:     "cyan",
			condition: true,
		},
		{
			symbol:    "*",
			logtype:   "INFO",
			color:     "cyan",
			condition: true,
		},
		{
			symbol:    "!",
			logtype:   "WARN",
			color:     "yellow",
			condition: true,
		},
		{
			symbol:    "X",
			logtype:   "ERROR",
			color:     "red",
			condition: true,
		},
		{
			symbol:    "+",
			logtype:   "DONE",
			color:     "green",
			condition: true,
		},
		{
			symbol:    ">",
			logtype:   "START",
			color:     "green",
			condition: true,
		},
	}

	for _, tc := range tests {
		old := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		switch tc.logtype {
		case "DEBUG":
			logger.Debug(module, target, message)
		case "INFO":
			logger.Info(module, target, message)
		case "WARN":
			logger.Warn(module, target, message)
		case "ERROR":
			logger.Error(module, target, message)
		case "DONE":
			logger.Done(module, target, message)
		case "START":
			logger.Start(module, target, message)
		}
		w.Close()
		out, _ := io.ReadAll(r)
		os.Stdout = old

		if string(out) == "" {
			t.Errorf("Unexpected output: got empty, want non-empty")
		}

		mod := EnsureLength(module, 35)
		logtype := EnsureLength(tc.logtype, 5)
		targ := EnsureLength(target, 22)

		expectedOutput := fmt.Sprintf(colors.Color("["+tc.color+"][%s] %s | [magenta]%s | [green]%s[reset]\t%s\n"), tc.symbol, logtype, targ, mod, message)

		if string(out) != expectedOutput {
			t.Errorf("Unexpected output: got %q, want %q", string(out), expectedOutput)
		}
	}
}

func capturePrintOutput(symbol, logtype, module, target, message, color string, condition bool) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	l := ILogger{Config: config.Config{Debug: condition}}
	l.Printer(symbol, logtype, module, target, message, color, condition)

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = old

	return string(out)
}

func TestLogger(t *testing.T) {
	nonNilConfig := &config.Config{Debug: true}
	nilConfig := (*config.Config)(nil)

	// Testing the function with a non-nil config
	result := Logger(nonNilConfig)
	if !result.Config.Debug {
		t.Errorf("Expected log level to be 'debug', but got %t", result.Config.Debug)
	}

	// Testing the function with a nil config
	result = Logger(nilConfig)
	if result.Config.Debug {
		t.Errorf("Expected default log level, but got %t", result.Config.Debug)
	}
}

func TestTicker(t *testing.T) {
	l := ILogger{}
	dt := func() time.Time { return time.Date(2024, 3, 30, 12, 0, 0, 0, time.UTC) }
	dtFormatted := dt().Format("15:04:05")

	testCases := []struct {
		tasks    map[string]bool
		running  int
		expected string
	}{
		{
			tasks:    make(map[string]bool),
			running:  0,
			expected: colors.Color("[green][*] " + dtFormatted + " | [magenta]test_target | [reset]Still running [yellow]0[reset] Tasks\n[yellow][reset]\n"),
		},
		{
			tasks: map[string]bool{
				"task1": true,
			},
			running:  1,
			expected: colors.Color("[green][*] " + dtFormatted + " | [magenta]test_target | [reset]Still running [yellow]1[reset] Tasks\n[yellow]task1[reset]\n"),
		},
	}

	for _, tc := range testCases {
		ActiveTasks = tc.tasks
		RunningTasks = tc.running

		old := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		l.Ticker("test_target", dt)

		w.Close()
		out, _ := io.ReadAll(r)
		os.Stdout = old

		got := string(out)
		if got != tc.expected {
			t.Errorf("Unexpected output, got: %s, want: %s", got, tc.expected)
		}
	}
}
