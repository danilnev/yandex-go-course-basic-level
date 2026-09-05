package main

import (
	"fmt"
)

func main() {
	var a, b, c float64

	fmt.Scanln(&a, &b, &c)

	if a == b && b == c && a == c {
		fmt.Println("Максимальное равенство")
	} else {
		fmt.Println("Не равны")
	}
}
