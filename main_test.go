package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	result := Hello()
	expected := "hello"

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	} else {
		t.Logf("Hello() PASSED. Expected %s, got %s", "Hello", result)
	}
}
