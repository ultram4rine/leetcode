package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{-1, 0}, -1))
}

func twoSum(numbers []int, target int) []int {
	var res = make([]int, 0, 2)

loop:
	for i := range numbers {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] > target {
				break
			} else if numbers[i]+numbers[j] == target {
				res = append(res, i+1, j+1)
				break loop
			}
		}
	}

	return res
}
