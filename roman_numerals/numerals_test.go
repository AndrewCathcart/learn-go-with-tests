package main

import "testing"

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Description string
		Arabic      int
		Expected    string
	}{
		{"1 gets converted to I", 1, "I"},
		{"2 gets converted to II", 2, "II"}
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			actual := ConvertToRoman(test.Arabic)
			if actual != test.Expected {
				t.Errorf("got %q, want %q", got, test.Want)
		}
		})
	}
}
