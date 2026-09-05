package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64

	fmt.Scanln(&num)

	if num < 0 {
		fmt.Println(-1)
		return
	}

	result := math.Sqrt(num)

	fmt.Printf("%.3f\n", result)
}
