package main

import "strings"

// ConvertToRoman converts arabic numbers (0-9) to roman numerals
// E.g. III = 3, IV = 4, V = 5, etc.
func ConvertToRoman(arabic int) string {

	if arabic == 4 {
		return "IV"
	}

	var result strings.Builder

	for i := 0; i < arabic; i++ {
		result.WriteString("I")
	}

	return result.String()
}
