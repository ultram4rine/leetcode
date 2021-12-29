package main

func main() {}

func searchMatrix(matrix [][]int, target int) bool {
	for _, r := range matrix {
		if target >= r[0] && target <= r[len(r)-1] {
			for _, n := range r {
				if target == n {
					return true
				}
			}
		}
	}

	return false
}
