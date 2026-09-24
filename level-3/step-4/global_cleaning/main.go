package main

import (
	"fmt"
	"unicode/utf8"
)

var (
	minimalLenthOfKey = 6
)

func DeleteLongKeys(m map[string]int) map[string]int {
	resultMap := map[string]int{}

	for key, value := range m {
		if utf8.RuneCountInString(key) >= minimalLenthOfKey {
			resultMap[key] = value
		}
	}

	return resultMap
}

func main() {
	fmt.Println(DeleteLongKeys(map[string]int{"danil_nevinsky": 5563, "danil": 46, "sanya_kdjfjka": 4336}))
}
