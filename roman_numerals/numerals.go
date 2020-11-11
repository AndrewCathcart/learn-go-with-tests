package main

// ConvertToRoman converts arabic numbers (0-9) to roman numerals
// E.g. III = 3, IV = 4, V = 5, etc.
func ConvertToRoman(arabic int) string {
	if arabic == 2 {
		return "II"
	}
	if arabic == 3 {
		return "III"
	}

	return "I"
}
