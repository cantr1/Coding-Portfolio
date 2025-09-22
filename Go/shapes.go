package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width, Height float64
}

func (c Circle) Area() float64 {
	return (3.14 * math.Pow(c.Radius, 2))
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func (r Rectangle) Area() float64 {
	return float64(r.Height * r.Width)
}

func (r Rectangle) Perimeter() float64 {
	return float64((2 * r.Height) + (2 * r.Width))
}

func PrintShapeInfo(s Shape) {
	fmt.Printf("Shape Type: %T\n", s)
	fmt.Printf("Shape Perimiter: %.2f\n", s.Perimeter())
	fmt.Printf("Shape Area: %.2f\n", s.Area())
}

func main() {
	my_circle := Circle{Radius: 5}
	PrintShapeInfo(my_circle)
}
