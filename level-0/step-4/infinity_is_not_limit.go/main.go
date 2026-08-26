package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64

	fmt.Scanln(&a)
	fmt.Scanln(&b)

	result := math.Round(math.Max(a, b))
	fmt.Println(result)
}
