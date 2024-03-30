package logger

import (
	"fmt"
	"gorecon/config"
	"io"
	"os"
	"testing"
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
		out := capturePrintOutput(tc.symbol, tc.logtype, module, target, message, tc.color, tc.condition)
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
