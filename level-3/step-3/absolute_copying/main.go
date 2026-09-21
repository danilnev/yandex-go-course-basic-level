package main

import "fmt"

func SliceCopy(nums []int) []int {
	result_slice := make([]int, len(nums))

	copy(result_slice, nums)

	return result_slice
}

func main() {
	slice := make([]int, 3, 6)
	slice[0] = 0
	slice[1] = 5
	slice[2] = 12
	fmt.Println(slice)
	new_slice := SliceCopy(slice)
	fmt.Println(new_slice)
	fmt.Println(cap(slice))
	fmt.Println(cap(new_slice))
}
