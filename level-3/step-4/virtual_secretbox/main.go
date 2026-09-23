package main

import (
	"fmt"
)

func FindMaxKey(m map[int]int) int {
	var maxKey int
	found := false

	for key := range m {
		if !found || key > maxKey {
			maxKey = key
			found = true
		}
	}

	return maxKey
}

func main() {
	fmt.Println(FindMaxKey(map[int]int{10: 39, 88: 43}))
}
