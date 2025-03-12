package main

import (
	"testing"
)

func TestGetMessage(t *testing.T) {
	got := getMessage()
	if len(got) == 0 {
		t.Errorf("getMessage() returned an empty string, want non-empty string")
	}
}

func TestGetMessageDoesNotModifyOriginalString(t *testing.T) {
	original := "Hello, Pipe!"
	expected := original

	got := getMessage()

	if got != expected {
		t.Errorf("getMessage() modified the original string, got: %s, want: %s", got, expected)
	}
}
