package plugins

import (
	"reflect"
	"testing"
)

func TestDirbuster(t *testing.T) {
	expectedModuleName := "dirbuster"
	expectedTags := []string{"default", "http"}
	expectedCommand := "feroxbuster"
	expectedArguments := []string{"-u", "{{.TargetPos}}", "-v", "-k", "-q", "-r", "-e", "-o", "{{.OutputFile}}"}
	expectedTargetFormat := "{{.Scheme}}://{{.Target}}:{{.Port}}"
	expectedTargetInplace := true
	expectedMatchPattern := "^http"
	expectedOutputFormat := "results/{{.Target}}/scans/{{.Port}}-{{.Protocol}}/dirbuster.txt"

	result := Dirbuster()

	if result.Name != expectedModuleName {
		t.Errorf("Expected Name to be %s, but got %s", expectedModuleName, result.Name)
	}
	if !sliceEqual(result.Tags, expectedTags) {
		t.Errorf("Expected Tags to be %v, but got %v", expectedTags, result.Tags)
	}
	if result.Command != expectedCommand {
		t.Errorf("Expected Command to be %s, but got %s", expectedCommand, result.Command)
	}
	if !sliceEqual(result.Arguments, expectedArguments) {
		t.Errorf("Expected Arguments to be %v, but got %v", expectedArguments, result.Arguments)
	}
	if result.TargetFormat != expectedTargetFormat {
		t.Errorf("Expected TargetFormat to be %s, but got %s", expectedTargetFormat, result.TargetFormat)
	}
	if result.TargetInplace != expectedTargetInplace {
		t.Errorf("Expected TargetInplace to be %t, but got %t", expectedTargetInplace, result.TargetInplace)
	}
	if result.MatchPattern != expectedMatchPattern {
		t.Errorf("Expected MatchPattern to be %s, but got %s", expectedMatchPattern, result.MatchPattern)
	}
	if result.OutputFormat != expectedOutputFormat {
		t.Errorf("Expected OutputFormat to be %s, but got %s", expectedOutputFormat, result.OutputFormat)
	}
}

func sliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestEnum4Linux(t *testing.T) {
	// Testing Name field
	expectedName := "enum4linux"
	result := Enum4Linux()
	if result.Name != expectedName {
		t.Errorf("Expected Name to be %s, but got %s", expectedName, result.Name)
	}

	// Testing Description field
	expectedDescription := "Enum4Linux Samba Enumeration"
	if result.Description != expectedDescription {
		t.Errorf("Expected Description to be %s, but got %s", expectedDescription, result.Description)
	}

	// Testing Tags field
	expectedTags := []string{"default", "smb"}
	if !reflect.DeepEqual(result.Tags, expectedTags) {
		t.Errorf("Expected Tags to be %v, but got %v", expectedTags, result.Tags)
	}

	// Testing Command field
	expectedCommand := "enum4linux"
	if result.Command != expectedCommand {
		t.Errorf("Expected Command to be %s, but got %s", expectedCommand, result.Command)
	}

	// Testing Arguments field
	expectedArguments := []string{"-a"}
	if !reflect.DeepEqual(result.Arguments, expectedArguments) {
		t.Errorf("Expected Arguments to be %v, but got %v", expectedArguments, result.Arguments)
	}

	// Testing ArgumentsInPlace field
	expectedArgumentsInPlace := true
	if result.ArgumentsInPlace != expectedArgumentsInPlace {
		t.Errorf("Expected ArgumentsInPlace to be %t, but got %t", expectedArgumentsInPlace, result.ArgumentsInPlace)
	}

	// Testing TargetAppend field
	expectedTargetAppend := true
	if result.TargetAppend != expectedTargetAppend {
		t.Errorf("Expected TargetAppend to be %t, but got %t", expectedTargetAppend, result.TargetAppend)
	}

	// Testing MatchPattern field
	expectedMatchPattern := "^(netbios-ssn|microsoft-ds|ldap|smb)"
	if result.MatchPattern != expectedMatchPattern {
		t.Errorf("Expected MatchPattern to be %s, but got %s", expectedMatchPattern, result.MatchPattern)
	}

	// Testing OutputFormat field
	expectedOutputFormat := "results/{{.Target}}/scans/{{.Port}}-{{.Protocol}}/enum4linux.txt"
	if result.OutputFormat != expectedOutputFormat {
		t.Errorf("Expected OutputFormat to be %s, but got %s", expectedOutputFormat, result.OutputFormat)
	}

	// Testing TargetFormat field
	expectedTargetFormat := "{{.Target}}"
	if result.TargetFormat != expectedTargetFormat {
		t.Errorf("Expected TargetFormat to be %s, but got %s", expectedTargetFormat, result.TargetFormat)
	}

	// Testing OutFile field
	expectedOutFile := true
	if result.OutFile != expectedOutFile {
		t.Errorf("Expected OutFile to be %t, but got %t", expectedOutFile, result.OutFile)
	}
}
