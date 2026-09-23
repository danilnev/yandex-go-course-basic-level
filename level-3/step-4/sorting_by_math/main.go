package main

import "fmt"

func CountingSort(contacts []string) map[string]int {
	resultMap := make(map[string]int)

	for _, value := range contacts {
		resultMap[value]++
	}

	return resultMap
}

func main() {
	fmt.Println(CountingSort([]string{"g", "g", "f", "g", "fg", "f"}))
}
