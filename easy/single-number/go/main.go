package main

func main() {}

func singleNumber(nums []int) int {
	var res = 0
	for _, n := range nums {
		res ^= n
	}
	return res
}
