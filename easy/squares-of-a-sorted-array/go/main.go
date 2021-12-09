package main

import "fmt"

func main() {
	fmt.Println(sortedSquares([]int{-7, -3, 2, 3, 11}))
}

func sortedSquares(nums []int) []int {
	var (
		answer = make([]int, len(nums))
		left   = 0
		right  = len(nums) - 1
		idx    = right
	)
	for left <= right {
		leftVal := nums[left] * nums[left]
		rightVal := nums[right] * nums[right]

		if leftVal >= rightVal {
			left++
			answer[idx] = leftVal
		} else {
			right--
			answer[idx] = rightVal
		}

		idx--
	}
	return answer
}
