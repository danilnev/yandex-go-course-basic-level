package main

import (
	"fmt"
)

func main() {
	var a, b int

	fmt.Scanln(&a, &b)

	if a > b {
		fmt.Println("Первое число больше второго")
		return
	} else if a < b {
		fmt.Println("Второе число больше первого")
		return
	}

	fmt.Println("Числа равны")
}
