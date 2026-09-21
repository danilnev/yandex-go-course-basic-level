package main

import "fmt"

func Join(nums1, nums2 []int) []int {
	result_slice := make([]int, len(nums1)+len(nums2))

	for i := 0; i < len(nums1)+len(nums2); i++ {
		if i < len(nums1) {
			result_slice[i] = nums1[i]
		} else {
			result_slice[i] = nums2[i-len(nums1)]
		}
	}

	return result_slice
}

func main() {
	slice1 := []int{}
	slice2 := []int{35, 546, 7, 3, 5, 6}
	fmt.Println(Join(slice1, slice2))
	fmt.Println(cap(Join(slice1, slice2)))
}
