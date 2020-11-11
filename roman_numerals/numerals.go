package main

import "strings"

// ConvertToRoman converts arabic numbers (0-9) to roman numerals
// E.g. III = 3, IV = 4, V = 5, etc.
func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for i := arabic; i > 0; i-- {
		if i == 4 {
			result.WriteString("IV")
			break
		}
		result.WriteString("I")
	}

	return result.String()
}
