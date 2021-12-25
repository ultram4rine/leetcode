package main

func main() {}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := range nums2 {
		for j := range nums1 {
			if nums2[i] < nums1[j] {
				copy(nums1[j+1:], nums1[j:])
				nums1[j] = nums2[i]
				m++
				break
			} else {
				if nums1[j] == 0 && j >= m {
					nums1[j] = nums2[i]
					m++
					break
				}
			}
		}
	}
}
