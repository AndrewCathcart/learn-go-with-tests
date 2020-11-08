package structs

import "testing"

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t *testing.T, shape Shape, expected float64) {
		t.Helper()
		actual := shape.Perimeter()
		if actual != expected {
			t.Errorf("got %g want %g", actual, expected)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		expected := 40.0
		checkPerimeter(t, rectangle, expected)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		expected := 62.83185307179586
		checkPerimeter(t, circle, expected)
	})
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, expected float64) {
		t.Helper()
		actual := shape.Area()
		if actual != expected {
			t.Errorf("got %g want %g", actual, expected)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		expected := 72.0
		checkArea(t, rectangle, expected)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		expected := 314.1592653589793
		checkArea(t, circle, expected)
	})
}
