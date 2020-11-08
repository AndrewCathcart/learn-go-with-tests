package structs

import "testing"

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		shape    Shape
		expected float64
	}{
		{Rectangle{10, 10}, 40.0},
		{Circle{10}, 62.83185307179586},
		{Triangle{1, 2, 3}, 6},
	}

	for _, tt := range perimeterTests {
		actual := tt.shape.Perimeter()
		if actual != tt.expected {
			t.Errorf("%#v got %g expected %g", tt.shape, actual, tt.expected)
		}
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, expected: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, expected: 314.1592653589793},
		{name: "Triangle", shape: Triangle{A: 3, B: 5, C: 4}, expected: 6},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.shape.Area()
			if actual != tt.expected {
				t.Errorf("%#v got %g expected %g", tt.shape, actual, tt.expected)
			}
		})
	}
}
