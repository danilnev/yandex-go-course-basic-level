package main

import "fmt"

func Mix(nums []int) []int {
	result_slice := make([]int, len(nums))
	nums1 := nums[:len(nums)/2]
	nums2 := nums[len(nums)/2:]

	fmt.Println(nums1, nums2)
	return result_slice
}

func main() {
	Mix([]int{1, 3, 5, 2, 4, 6})
}
