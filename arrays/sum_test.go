package sum

import "testing"

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3}

	actual := Sum(numbers)
	expected := 6

	if actual != expected {
		t.Errorf("actual %d expected %d given, %v", actual, expected, numbers)
	}
}
