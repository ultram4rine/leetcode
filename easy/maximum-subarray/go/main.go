package main

func main() {}

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	var (
		maxSoFar = nums[0]
		currMax  = nums[0]
	)

	for i := 1; i < len(nums); i++ {
		currMax = max(nums[i], currMax+nums[i])
		maxSoFar = max(maxSoFar, currMax)
	}

	return maxSoFar
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
