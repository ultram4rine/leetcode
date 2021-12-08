package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
}

func searchInsert(nums []int, target int) int {
	if target < nums[0] {
		return 0
	} else if target > nums[len(nums)-1] {
		return len(nums)
	}

	var (
		left   int = 0
		right  int = len(nums)
		midIdx     = (left + right) / 2
		mid        = nums[midIdx]
	)
	for left <= right {
		if target < mid {
			right = midIdx - 1
		} else if target > mid {
			left = midIdx + 1
		} else {
			return midIdx
		}

		midIdx = (left + right) / 2
		mid = nums[midIdx]
	}

	return midIdx + 1
}
