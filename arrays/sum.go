package sum

// Sum takes an array and returns the sum of all values
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
