package main

import (
	"fmt"
	"math"
)

func SquareRoots(a, b, c float64) (float64, float64) {
	discriminant := findDiscriminant(a, b, c)
	x1, x2 := 0.0, 0.0
	if discriminant > 0 {
		x2 = (-b + math.Sqrt(discriminant)) / (2 * a)
		x1 = (-b - math.Sqrt(discriminant)) / (2 * a)
	} else if discriminant == 0 {
		x1 = -b / (2 * a)
		x2 = x1
	}
	return x1, x2
}

func findDiscriminant(a, b, c float64) float64 {
	return math.Pow(b, 2) - 4*a*c
}

func main() {
	// var a, b, c float64
	// fmt.Scanln(&a, &b, &c)
	// fmt.Println(SquareRoots(a, b, c))

	fmt.Println(SquareRoots(1, -3, 2))
	fmt.Println(SquareRoots(1, 4, 4.0))
	fmt.Println(SquareRoots(1, 1, 1))
	fmt.Println(SquareRoots(4, 4, 1))
	fmt.Println(SquareRoots(4, 4, -1))
	fmt.Println(SquareRoots(2, 2, 8))
	fmt.Println(SquareRoots(1, 4, -5))
	fmt.Println(SquareRoots(1, 0, -9))
	fmt.Println(SquareRoots(1, 5, 0))
}
