package main

import (
	"fmt"
)

func GoOrNot(str string) {
	if str == "Go" {
		fmt.Println("Go!")
	} else {
		fmt.Println("Я знаю только Go!")
	}
}

func main() {
	var str string

	fmt.Scanln(&str)
	GoOrNot(str)
}
