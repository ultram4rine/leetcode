package main

import (
	"fmt"
)

func main() {
	fmt.Println(search([]int{-1, 0, 3, 4, 5, 9, 12}, 11))
}

func search(nums []int, target int) int {
	if target < nums[0] || target > nums[len(nums)-1] {
		return -1
	}

	var (
		left  int = 0
		right int = len(nums)
	)
	for left <= right {
		var (
			midIdx = (left + right) / 2
			mid    = nums[midIdx]
		)

		if target < mid {
			right = midIdx - 1
		} else if target > mid {
			left = midIdx + 1
		} else {
			return midIdx
		}
	}

	return -1
}
