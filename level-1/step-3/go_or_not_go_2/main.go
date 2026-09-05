package main

import (
	"fmt"
)

func main() {
	for range 5 {
		var str string
		fmt.Scanln(&str)
		if str == "Go" {
			fmt.Println("Go!")
			continue
		}
		fmt.Println("Я знаю только Go!")
	}
}
