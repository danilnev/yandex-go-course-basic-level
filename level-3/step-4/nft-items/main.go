package main

import (
	"fmt"
)

func SumOfValuesInMap(m map[int]int) int {
	var sum int

	for _, value := range m {
		sum += value
	}

	return sum
}

func main() {
	fmt.Println(SumOfValuesInMap(map[int]int{30: 49, 8: -1}))
}
