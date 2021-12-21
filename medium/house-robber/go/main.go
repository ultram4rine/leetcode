package main

func main() {}

func rob(nums []int) int {
	var n = len(nums)
	if n == 0 {
		return 0
	}

	var val1 = nums[0]
	if n == 1 {
		return val1
	}

	var val2 = maxInt(val1, nums[1])
	if n == 2 {
		return val2
	}

	var max = 0

	for i := 2; i < n; i++ {
		max = maxInt(nums[i]+val1, val2)
		val1 = val2
		val2 = max
	}

	return max
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
