package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestPort3000IsWritten(t *testing.T) {
	// Read the content of your go_hello.go file
	content, err := ioutil.ReadFile("path/to/go_hello.go") // Update the path to your go_hello.go file
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	// Check if the content contains the string ":3000"
	if !strings.Contains(string(content), ":3000") {
		t.Errorf("Port 3000 is not written in go_hello.go")
	}
}
