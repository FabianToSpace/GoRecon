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
	expectedDescription := "Enum4Linux Samba Enumeration"
	expectedTags := []string{"default", "smb"}
	expectedCommand := "enum4linux"
	expectedArguments := []string{"-a"}
	expectedArgumentsInPlace := true
	expectedTargetAppend := true
	expectedMatchPattern := "^(netbios-ssn|microsoft-ds|ldap|smb)"
	expectedOutputFormat := "results/{{.Target}}/scans/{{.Port}}-{{.Protocol}}/enum4linux.txt"
	expectedTargetFormat := "{{.Target}}"
	expectedOutFile := true

	result := Enum4Linux()
	if result.Name != expectedName {
		t.Errorf("Expected Name to be %s, but got %s", expectedName, result.Name)
	}

	if result.Description != expectedDescription {
		t.Errorf("Expected Description to be %s, but got %s", expectedDescription, result.Description)
	}

	if !reflect.DeepEqual(result.Tags, expectedTags) {
		t.Errorf("Expected Tags to be %v, but got %v", expectedTags, result.Tags)
	}

	if result.Command != expectedCommand {
		t.Errorf("Expected Command to be %s, but got %s", expectedCommand, result.Command)
	}

	if !reflect.DeepEqual(result.Arguments, expectedArguments) {
		t.Errorf("Expected Arguments to be %v, but got %v", expectedArguments, result.Arguments)
	}

	if result.ArgumentsInPlace != expectedArgumentsInPlace {
		t.Errorf("Expected ArgumentsInPlace to be %t, but got %t", expectedArgumentsInPlace, result.ArgumentsInPlace)
	}

	if result.TargetAppend != expectedTargetAppend {
		t.Errorf("Expected TargetAppend to be %t, but got %t", expectedTargetAppend, result.TargetAppend)
	}

	if result.MatchPattern != expectedMatchPattern {
		t.Errorf("Expected MatchPattern to be %s, but got %s", expectedMatchPattern, result.MatchPattern)
	}

	if result.OutputFormat != expectedOutputFormat {
		t.Errorf("Expected OutputFormat to be %s, but got %s", expectedOutputFormat, result.OutputFormat)
	}

	if result.TargetFormat != expectedTargetFormat {
		t.Errorf("Expected TargetFormat to be %s, but got %s", expectedTargetFormat, result.TargetFormat)
	}

	if result.OutFile != expectedOutFile {
		t.Errorf("Expected OutFile to be %t, but got %t", expectedOutFile, result.OutFile)
	}
}

func TestNikto(t *testing.T) {
	expectedName := "nikto"
	expectedDescription := "Nikto Web Server Scanner"
	expectedTags := []string{"default", "http"}
	expectedCommand := "nikto"
	expectedArguments := []string{"-host", "{{.TargetPos}}", "-ask=no", "-Tuning=x4567890ac", "-nointeractive", "-r", "-e", "-o", "{{.OutputFile}}"}
	expectedTargetFormat := "{{.Scheme}}://{{.Target}}:{{.Port}}"
	expectedTargetInplace := true
	expectedMatchPattern := "^http"
	expectedOutputFormat := "results/{{.Target}}/scans/{{.Port}}-{{.Protocol}}/nikto.txt"
	expectedOutFile := true

	ss := Nikto()

	if ss.Name != expectedName {
		t.Errorf("Expected Name to be %s, got %s", expectedName, ss.Name)
	}
	if ss.Description != expectedDescription {
		t.Errorf("Expected Description to be %s, got %s", expectedDescription, ss.Description)
	}
	if !reflect.DeepEqual(ss.Tags, expectedTags) {
		t.Errorf("Expected Tags to be %v, got %v", expectedTags, ss.Tags)
	}
	if ss.Command != expectedCommand {
		t.Errorf("Expected Command to be %s, got %s", expectedCommand, ss.Command)
	}
	if !reflect.DeepEqual(ss.Arguments, expectedArguments) {
		t.Errorf("Expected Arguments to be %v, got %v", expectedArguments, ss.Arguments)
	}
	if ss.TargetFormat != expectedTargetFormat {
		t.Errorf("Expected TargetFormat to be %s, got %s", expectedTargetFormat, ss.TargetFormat)
	}
	if ss.TargetInplace != expectedTargetInplace {
		t.Errorf("Expected TargetInplace to be %t, got %t", expectedTargetInplace, ss.TargetInplace)
	}
	if ss.MatchPattern != expectedMatchPattern {
		t.Errorf("Expected MatchPattern to be %s, got %s", expectedMatchPattern, ss.MatchPattern)
	}
	if ss.OutputFormat != expectedOutputFormat {
		t.Errorf("Expected OutputFormat to be %s, got %s", expectedOutputFormat, ss.OutputFormat)
	}
	if ss.OutFile != expectedOutFile {
		t.Errorf("Expected OutFile to be %t, got %t", expectedOutFile, ss.OutFile)
	}
}

func TestNmapFtp(t *testing.T) {
	ns := NmapFtp()
	expectedName := "nmap-ftp"
	expectedDescription := "Nmap FTP Scripts"
	expectedTags := []string{"default", "ftp"}
	expectedCommand := "nmap"
	expectedArguments := []string{"-sV", "-p{{.Port}}", "-oN", "{{.OutputFile}}", "--script", "banner,(ftp* or ssl*) and safe"}
	expectedArgumentsInPlace := true
	expectedTargetAppend := true
	expectedMatchPattern := "^ftp"
	expectedOutputFormat := "results/{{.Target}}/scans/{{.Port}}-{{.Protocol}}/nmap-ftp.txt"
	expectedTargetFormat := "{{.Target}}"

	if ns.Name != expectedName {
		t.Errorf("Expected Name to be %s, but got %s", expectedName, ns.Name)
	}

	if ns.Description != expectedDescription {
		t.Errorf("Expected Description to be %s, but got %s", expectedDescription, ns.Description)
	}

	if !reflect.DeepEqual(ns.Tags, expectedTags) {
		t.Errorf("Expected Tags to be %v, but got %v", expectedTags, ns.Tags)
	}

	if ns.Command != expectedCommand {
		t.Errorf("Expected Command to be %s, but got %s", expectedCommand, ns.Command)
	}

	if !reflect.DeepEqual(ns.Arguments, expectedArguments) {
		t.Errorf("Expected Arguments to be %v, but got %v", expectedArguments, ns.Arguments)
	}

	if ns.ArgumentsInPlace != expectedArgumentsInPlace {
		t.Errorf("Expected ArgumentsInPlace to be %v, but got %v", expectedArgumentsInPlace, ns.ArgumentsInPlace)
	}

	if ns.TargetAppend != expectedTargetAppend {
		t.Errorf("Expected TargetAppend to be %v, but got %v", expectedTargetAppend, ns.TargetAppend)
	}

	if ns.MatchPattern != expectedMatchPattern {
		t.Errorf("Expected MatchPattern to be %s, but got %s", expectedMatchPattern, ns.MatchPattern)
	}

	if ns.OutputFormat != expectedOutputFormat {
		t.Errorf("Expected OutputFormat to be %s, but got %s", expectedOutputFormat, ns.OutputFormat)
	}

	if ns.TargetFormat != expectedTargetFormat {
		t.Errorf("Expected TargetFormat to be %s, but got %s", expectedTargetFormat, ns.TargetFormat)
	}
}

func TestNmapSmb(t *testing.T) {
	expectedModuleName := "nmap-smb"
	expectedDescription := "Nmap SMB Scripts"
	expectedTags := []string{"default", "smb"}
	expectedCommand := "nmap"
	expectedArguments := []string{"-sV", "-p{{.Port}}", "-oN", "{{.OutputFile}}", "--script", "smb-enum-* and safe"}
	expectedArgumentsInPlace := true
	expectedTargetAppend := true
	expectedMatchPattern := "^(netbios-ssn|microsoft-ds|ldap|smb)"
	expectedOutputFormat := "results/{{.Target}}/scans/{{.Port}}-{{.Protocol}}/" + expectedModuleName + ".txt"
	expectedTargetFormat := "{{.Target}}"

	scan := NmapSmb()

	if scan.Name != expectedModuleName {
		t.Errorf("Expected module name %s, but got %s", expectedModuleName, scan.Name)
	}

	if scan.Description != expectedDescription {
		t.Errorf("Expected description %s, but got %s", expectedDescription, scan.Description)
	}

	if !reflect.DeepEqual(scan.Tags, expectedTags) {
		t.Errorf("Expected tags %v, but got %v", expectedTags, scan.Tags)
	}

	if scan.Command != expectedCommand {
		t.Errorf("Expected command %s, but got %s", expectedCommand, scan.Command)
	}

	if !reflect.DeepEqual(scan.Arguments, expectedArguments) {
		t.Errorf("Expected arguments %v, but got %v", expectedArguments, scan.Arguments)
	}

	if scan.ArgumentsInPlace != expectedArgumentsInPlace {
		t.Errorf("Expected arguments in place %t, but got %t", expectedArgumentsInPlace, scan.ArgumentsInPlace)
	}

	if scan.TargetAppend != expectedTargetAppend {
		t.Errorf("Expected target append %t, but got %t", expectedTargetAppend, scan.TargetAppend)
	}

	if scan.MatchPattern != expectedMatchPattern {
		t.Errorf("Expected match pattern %s, but got %s", expectedMatchPattern, scan.MatchPattern)
	}

	if scan.OutputFormat != expectedOutputFormat {
		t.Errorf("Expected output format %s, but got %s", expectedOutputFormat, scan.OutputFormat)
	}

	if scan.TargetFormat != expectedTargetFormat {
		t.Errorf("Expected target format %s, but got %s", expectedTargetFormat, scan.TargetFormat)
	}
}

func TestWhatweb(t *testing.T) {
	expectedModuleName := "whatweb"
	expectedDescription := "Whatweb"
	expectedTags := []string{"default", "http"}
	expectedCommand := "whatweb"
	expectedArguments := []string{"--color=never", "--no-errors", "-a 3", "-v"}
	expectedTargetFormat := "{{.Scheme}}://{{.Target}}:{{.Port}}"
	expectedTargetAppend := true
	expectedMatchPattern := "^http"
	expectedOutputFormat := "results/{{.Target}}/scans/{{.Port}}-{{.Protocol}}/" + expectedModuleName + ".txt"
	expectedOutFile := true

	scan := Whatweb()

	if scan.Name != expectedModuleName {
		t.Errorf("Expected module name %s, but got %s", expectedModuleName, scan.Name)
	}

	if scan.Description != expectedDescription {
		t.Errorf("Expected description %s, but got %s", expectedDescription, scan.Description)
	}

	if !reflect.DeepEqual(scan.Tags, expectedTags) {
		t.Errorf("Expected tags %v, but got %v", expectedTags, scan.Tags)
	}

	if scan.Command != expectedCommand {
		t.Errorf("Expected command %s, but got %s", expectedCommand, scan.Command)
	}

	if !reflect.DeepEqual(scan.Arguments, expectedArguments) {
		t.Errorf("Expected arguments %v, but got %v", expectedArguments, scan.Arguments)
	}

	if scan.TargetFormat != expectedTargetFormat {
		t.Errorf("Expected target format %s, but got %s", expectedTargetFormat, scan.TargetFormat)
	}

	if scan.TargetAppend != expectedTargetAppend {
		t.Errorf("Expected target append %t, but got %t", expectedTargetAppend, scan.TargetAppend)
	}

	if scan.MatchPattern != expectedMatchPattern {
		t.Errorf("Expected match pattern %s, but got %s", expectedMatchPattern, scan.MatchPattern)
	}

	if scan.OutputFormat != expectedOutputFormat {
		t.Errorf("Expected output format %s, but got %s", expectedOutputFormat, scan.OutputFormat)
	}

	if scan.OutFile != expectedOutFile {
		t.Errorf("Expected out file %t, but got %t", expectedOutFile, scan.OutFile)
	}
}
