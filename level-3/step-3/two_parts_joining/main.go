package main

import "fmt"

func Mix(nums []int) []int {
	result_slice := make([]int, 0)
	nums1 := nums[:len(nums)/2]
	nums2 := nums[len(nums)/2:]

	for i := range len(nums) / 2 {
		result_slice = append(result_slice, nums1[i], nums2[i])
	}

	return result_slice
}

func main() {
	fmt.Println(Mix([]int{1, 3, 5, 2, 4, 6}))
}
