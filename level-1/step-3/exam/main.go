package main

import (
	"fmt"
)

func main() {
	var N int

	fmt.Scanln(&N)

	for range N {
		var mark float64
		fmt.Scanln(&mark)

		if 90 <= mark && mark <= 100 {
			fmt.Println("5")
		} else if 75 <= mark && mark <= 89 {
			fmt.Println("4")
		} else if 50 <= mark && mark <= 74 {
			fmt.Println("3")
		} else if 0 <= mark && mark <= 49 {
			fmt.Println("2")
		} else {
			fmt.Println("Неверный балл")
		}
	}
}
