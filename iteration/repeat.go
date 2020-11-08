package iteration

// Repeat takes a base string and concatenates it numRepetitions times
func Repeat(base string, numRepetitions int) string {
	var repeated string
	for i := 0; i < numRepetitions; i++ {
		repeated += base
	}
	return repeated
}
