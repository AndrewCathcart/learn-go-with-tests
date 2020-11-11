package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Andy")

	actual := buffer.String()
	expected := "Hello, Andy"

	if expected != actual {
		t.Errorf("got %q expected %q", actual, expected)
	}
}
