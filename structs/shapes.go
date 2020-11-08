package structs

import "math"

// Shape has an Area
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle has a Width and a Height
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter calculates the perimeter of a rectangle given a width and height
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Area calculates the area of a Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle has a Radius
type Circle struct {
	Radius float64
}

// Perimeter calculates the circumference of a Circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Area calculates the area of a C
func (c Circle) Area() float64 {
	return math.Pi * (math.Pow(c.Radius, 2.0))
}

// Triangle has a, b, c
type Triangle struct {
	A, B, C float64
}

// Perimeter calculates the perimeter of a Triangle
func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

// Area calculates the area of a Triangle (using Heron's Formula)
func (t Triangle) Area() float64 {
	s := (t.Perimeter() / 2)
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}
