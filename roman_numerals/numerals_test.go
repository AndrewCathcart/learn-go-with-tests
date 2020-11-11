package main

import "testing"

func TestRomanNumerals(t *testing.T) {
	actual := ConvertToRoman(1)
	expected := "I"

	if actual != expected {
		t.Errorf("got %q, expected %q", actual, expected)
	}
}
