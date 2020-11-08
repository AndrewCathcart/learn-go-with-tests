package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, actual string, expected string) {
		t.Helper() // helps indicate line numbers in failing tests
		if actual != expected {
			t.Errorf("got %q expected %q", actual, expected)
		}
	}

	t.Run("saying hello with a name", func(t *testing.T) {
		actual := Hello("Andy")
		expected := "Hello, Andy!"
		assertCorrectMessage(t, actual, expected)
	})

	t.Run("say 'Hello, World!' when no name is supplied", func(t *testing.T) {
		actual := Hello("")
		expected := "Hello, World!"
		assertCorrectMessage(t, actual, expected)
	})
}
