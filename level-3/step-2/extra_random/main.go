package main

func FindMaxMinInArray(array [10]int) (int, int) {
	min := array[0]
	max := array[0]
	for _, num := range array {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	return max, min
}
