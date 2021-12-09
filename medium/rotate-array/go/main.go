package main

import "fmt"

func main() {
	var nums = []int{-1, -100, 3, 99}
	rotate(nums, 2)
	fmt.Println(nums)
}

func rotate(nums []int, k int) {
	var (
		n   = len(nums)
		res = make([]int, len(nums))
	)
	copy(res, nums)
	for i := range nums {
		res[(i+k)%n] = nums[i]
	}
	copy(nums, res)
}
