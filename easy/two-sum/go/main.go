package main

func main() {}

func twoSum(nums []int, target int) []int {
	var m = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if j, ok := m[target-nums[i]]; ok {
			return []int{i, j}
		}
		m[nums[i]] = i
	}
	return []int{0, 0}
}
