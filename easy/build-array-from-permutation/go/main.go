package main

func main() {}

func buildArray(nums []int) []int {
	var ans = make([]int, 0)

	for i := 0; i < len(nums); i++ {
		ans = append(ans, nums[nums[i]])
	}

	return ans
}
