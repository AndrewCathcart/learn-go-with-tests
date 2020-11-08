package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3}

	actual := Sum(numbers)
	expected := 6

	if actual != expected {
		t.Errorf("actual %d expected %d given, %v", actual, expected, numbers)
	}
}

func TestSliceTotals(t *testing.T) {
	actual := SliceTotals([]int{1, 2}, []int{3, 4})
	expected := []int{3, 7}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v want %v", actual, expected)
	}
}

func TestSliceTailTotals(t *testing.T) {
	checkSums := func(t *testing.T, actual, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("got %v want %v", actual, expected)
		}
	}

	t.Run("sum tails of slices", func(t *testing.T) {
		actual := SliceTailTotals([]int{1, 2}, []int{3, 4})
		expected := []int{2, 4}
		checkSums(t, actual, expected)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		actual := SliceTailTotals([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}
		checkSums(t, actual, expected)
	})
}
