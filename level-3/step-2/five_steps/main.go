package main

import (
	"fmt"
)

func FiveSteps(array [5]int) [5]int {
	var result [5]int
	for i := 0; i < 5; i++ {
		result[i] = array[4-i]
	}
	return result
}

func main() {
	fmt.Println(FiveSteps([5]int{1, 2, 3, 4, 5}))
}
