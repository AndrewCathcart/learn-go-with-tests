package sum

// Sum takes an array and returns the sum of all values
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SliceTotals takes a variable number of slices, and returns a new slice containing the totals for each provided slice
func SliceTotals(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

// SliceTailTotals takes a variable number of slices, and returns a new slice containing the tail totals (everything except the first value) for each provided slice
func SliceTailTotals(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(numbers[1:]))
		}
	}

	return sums
}
