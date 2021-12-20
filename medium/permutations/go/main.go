package main

import "sort"

func main() {}

func permute(nums []int) [][]int {
	var (
		res = make([][]int, 0)
		fst = make([]int, len(nums))
	)

	sort.Ints(nums)
	copy(fst, nums)
	res = append(res, fst)

	for getNext(nums) {
		var curr = make([]int, len(nums))
		copy(curr, nums)
		res = append(res, curr)
	}

	return res
}

func getNext(nums []int) bool {
	for {
		var (
			n = len(nums)
			j = n - 2
		)

		for j != -1 && nums[j] >= nums[j+1] {
			j--
		}

		if j == -1 {
			return false
		}

		var m = n - 1
		for nums[j] >= nums[m] {
			m--
		}

		nums[j], nums[m] = nums[m], nums[j]

		var (
			l = j + 1
			r = n - 1
		)

		for l < r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}

		if !(j > n-1) {
			break
		}
	}

	return true
}
