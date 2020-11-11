package main

import "testing"

// Roman numerals kata - convert arabic number (0-9) to a Roman Numeral
// E.g. III = 3, IV = 4, V = 5, etc.
func TestRomanNumerals(t *testing.T) {
	actual := ConvertToRoman(1)
	expected := "I"

	if actual != expected {
		t.Errorf("got %q, expected %q", actual, expected)
	}
}
