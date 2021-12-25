package main

func main() {}

func intersect(nums1 []int, nums2 []int) []int {
	var (
		res = make([]int, 0)
		m   = make(map[int]int)
	)

	for i := range nums1 {
		m[nums1[i]] += 1
	}

	for j := range nums2 {
		if k, ok := m[nums2[j]]; k != 0 && ok {
			res = append(res, nums2[j])
			m[nums2[j]] -= 1
		}
	}

	return res
}
