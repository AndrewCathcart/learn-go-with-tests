package structs

import "testing"

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		shape    Shape
		expected float64
	}{
		{Rectangle{10, 10}, 40.0},
		{Circle{10}, 62.83185307179586},
	}

	for _, tt := range perimeterTests {
		actual := tt.shape.Perimeter()
		if actual != tt.expected {
			t.Errorf("got %g expected %g", actual, tt.expected)
		}
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape    Shape
		expected float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		actual := tt.shape.Area()
		if actual != tt.expected {
			t.Errorf("got %g expected %g", actual, tt.expected)
		}
	}
}
