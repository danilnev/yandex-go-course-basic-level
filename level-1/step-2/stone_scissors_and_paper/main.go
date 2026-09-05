package main

import (
	"fmt"
)

func main() {
	var first, second string

	fmt.Scanln(&first)
	fmt.Scanln(&second)

	if first == second {
		fmt.Println("Ничья")
	} else if (first == "камень" && second == "ножницы") || (first == "ножницы" && second == "бумага") || (first == "бумага" && second == "камень") {
		fmt.Println("Первый игрок победил")
	} else {
		fmt.Println("Второй игрок победил")
	}
}
