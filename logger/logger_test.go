package logger

import "testing"

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
