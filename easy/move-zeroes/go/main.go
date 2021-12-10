package main

import "fmt"

func main() {
	var nums = []int{0, 1, 0, 0, 3, 12, 0, 0, 0, 0}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	var k = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, 0)
			k++
			i--
		}
		if i > len(nums)-k {
			break
		}
	}
}
