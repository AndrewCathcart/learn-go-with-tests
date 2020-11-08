package main

import "testing"

func TestHello(t *testing.T) {
	actual := Hello("Andy")
	expected := "Hello, Andy!"

	if actual != expected {
		t.Errorf("got %q want %q", actual, expected)
	}
}
