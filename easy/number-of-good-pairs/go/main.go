package main

func main() {}

func numIdenticalPairs(nums []int) int {
	var ans int
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				ans++
			}
		}
	}
	return ans
}
