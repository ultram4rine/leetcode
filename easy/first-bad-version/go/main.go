package main

import "fmt"

func main() {
	fmt.Println(firstBadVersion(100))
}

func isBadVersion(version int) bool {
	fmt.Println(version)
	return version >= 20
}

func firstBadVersion(n int) int {
	if n == 1 {
		return 1
	}

	var (
		left  int = 0
		right int = n
		mid       = (left + right) / 2
	)
	for left <= right {
		if isBadVersion(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}

		mid = (left + right) / 2
	}

	return mid + 1
}
