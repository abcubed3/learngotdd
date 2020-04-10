package main

import (
	"fmt"
	"math"
)

//interface Shape
type Shape interface {
	Area() float64
	Perimeter() float64
}

//Rectangle Struct
type Rectangle struct {
	width  float64
	height float64
}

//Triangle Struct
type Triangle struct {
	base   float64
	height float64
}

// Method for Rectangle struct area
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Method for Triangle struct area
func (t Triangle) Area() float64 {
	return t.base * t.height * 0.5
}

//Perimeter function
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Cicle Struct
type Circle struct {
	radius float64
}

// Method for Rectangle struct area
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Method for Triangle struct perimeter
func (t Triangle) Perimeter() float64 {
	return t.base + math.Sqrt(t.height)
}
func main() {
	rectangle := Rectangle{10.0, 10.0}
	circle := Circle{10.0}
	triangle := Triangle{10, 2}
	fmt.Println(rectangle.Area())
	fmt.Println(circle.Area())
	fmt.Println(triangle.Area())
}
