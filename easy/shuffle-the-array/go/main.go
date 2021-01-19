package main

func main() {}

func shuffle(nums []int, n int) []int {
	answer := make([]int, 0, 2*n)
	for i := 0; i < n; i++ {
		answer = append(answer, nums[i])
		answer = append(answer, nums[n+i])
	}
	return answer
}
