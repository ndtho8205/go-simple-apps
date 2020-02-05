package structs

import "math"

type Shape interface {
	Perimeter() float64
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

type Square struct {
	Length float64
}

func (rectangle Rectangle) Perimeter() float64 {
	return (rectangle.Width + rectangle.Height) * 2
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.Width * rectangle.Height
}

func (circle Circle) Perimeter() float64 {
	return 2 * math.Pi * circle.Radius
}

func (circle Circle) Area() float64 {
	return math.Pi * circle.Radius * circle.Radius
}

func (square Square) Perimeter() float64 {
	return 4 * square.Length
}

func (square Square) Area() float64 {
	return square.Length * square.Length
}
