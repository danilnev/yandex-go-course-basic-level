package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)
	min := math.Min(a, math.Min(b, c))
	fmt.Println(min)
}
