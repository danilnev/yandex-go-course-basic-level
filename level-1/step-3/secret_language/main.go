package main

import (
	"fmt"
)

func main() {
	var word string

	fmt.Scanln(&word)
	result := ""

	for _, letter := range word {
		if letter == 'а' || letter == 'о' {
			continue
		}
		result += string(letter)
	}

	fmt.Println(result)
}
