package main

import (
	"fmt"
	"strings"
)

func main() {
	var N int
	var word, str string

	fmt.Scanln(&N)
	fmt.Scanln(&word)

	count := 0
	for i := 0; i < N; i++ {
		if i == N-1 {
			fmt.Scanln(&str)
		} else {
			fmt.Scan(&str)
		}
		if strings.EqualFold(str, word) {
			count++
		}
	}

	fmt.Println(count)
}
