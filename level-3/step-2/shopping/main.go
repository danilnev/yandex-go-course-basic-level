package main

func SumOfArray(array [6]int) int {
	sum := 0
	for _, item := range array {
		sum += item
	}
	return sum
}
