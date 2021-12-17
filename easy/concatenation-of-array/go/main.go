package main

func main() {}

func getConcatenation(nums []int) []int {
	var (
		n   = len(nums)
		ans = make([]int, 2*n)
	)

	for i := range nums {
		ans[i] = nums[i]
		ans[i+n] = nums[i]
	}

	return ans
}
