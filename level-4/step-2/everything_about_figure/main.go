package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func main() {
	var circle Shape = Circle{radius: 2.5}
	var rectangle Shape = Rectangle{width: 4.5, height: 5.945}
	fmt.Println(circle.Area(), rectangle.Area())
}
