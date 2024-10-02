package exercises

import "fmt"

// Write a program that creates two custom struct types called 'triangle' and 'square'

// The 'square' type should be a struct with a field called 'sideLength' of type float64
type square struct {
	sideLength float64
}

// The 'triangle' type should be a struct with a field called 'height' of type float64
// and a field of type 'base' of type float64
type triangle struct {
	height float64
	base   float64
}

// Add a 'shape' interface that defines a function called 'printArea'.
type shape interface {
	getArea() float64
}

func GeoMain() {
	fmt.Println("------- Geometry -------")
	t := triangle{
		10,
		10,
	}
	s := square{10}
	printArea(t)
	printArea(s)
}

// Design the interface so that the 'printArea' function can be called with either a triangle or a square.
// This function should calculate the area of the given shape and print it out to the terminal.
func printArea(s shape) float64 {
	area := s.getArea()
	fmt.Println(area)
	return area
}

// Both types should have a function called 'getArea' that returns the calculated area of the square or triangle
func (t triangle) getArea() float64 {
	// Area of triangle = 0.5 * base * height
	return (0.5 * t.base * t.height)
}

func (s square) getArea() float64 {
	// Area of square  = sideLength * sideLength
	return (s.sideLength * s.sideLength)
}
