package main

import (
	"fmt"
)

func main() {
	var N int
	var sale float64

	fmt.Scanln(&N)
	fmt.Scanln(&sale)

	sum := 0.0

	for range N {
		var price float64

		fmt.Scanln(&price)

		sum += price * ((100.0 - sale) / 100.0)
	}

	fmt.Println(sum)
}
