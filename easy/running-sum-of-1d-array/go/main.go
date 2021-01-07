package main

func main() {}

func runningSum(nums []int) []int {
	var answer []int

	for i := range nums {
		if i == 0 {
			answer = append(answer, nums[i])
		} else {
			answer = append(answer, answer[i-1]+nums[i])
		}
	}

	return answer
}
