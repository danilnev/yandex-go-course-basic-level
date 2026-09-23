package main

import (
	"errors"
	"fmt"
)

var (
	IncorrectNums = errors.New("Некорректный массив чисел")
	IncorrectN    = errors.New("Некорректная длина массива")
)

func UnderLimit(nums []int, limit int, n int) ([]int, error) {
	if nums == nil {
		return nil, IncorrectNums
	} else if n <= 0 {
		return nil, IncorrectN
	}

	result_slice := make([]int, 0)

	for i, item := range nums {
		if len(result_slice) == n {
			break
		}
		if item < limit {
			result_slice = append(result_slice, nums[i])
		}
	}

	return result_slice, nil
}

func main() {
	fmt.Println(UnderLimit([]int{4, 7, 89, 3, 21, 2, 5, 7, 32, 4, 6, 8, 0, 3, 4, 6, 2, 115, 12}, 3, 5))
}
