package main

import (
	"fmt"
)

func main() {
	var answer string

	fmt.Scanln(&answer)
	for answer != "да" && answer != "нет" && answer != "чёрный" && answer != "белый" {
		fmt.Println("Игра продолжается")
		fmt.Scanln(&answer)
	}

	fmt.Println("Поражение")
}
