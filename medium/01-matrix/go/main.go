package main

import (
	"fmt"
)

func main() {
	fmt.Println(updateMatrix([][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}))
}

func updateMatrix(mat [][]int) [][]int {
	var m = len(mat)
	if m == 0 {
		return mat
	}
	var n = len(mat[0])
	res := make([][]int, m)

	for i := range mat {
		res[i] = make([]int, n)
		for j := range mat[i] {
			res[i][j] = 10000
			if mat[i][j] == 0 {
				res[i][j] = 0
			} else {
				if i > 0 {
					res[i][j] = min(res[i][j], res[i-1][j]+1)
				}
				if j > 0 {
					res[i][j] = min(res[i][j], res[i][j-1]+1)
				}
			}
		}
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i < m-1 {
				res[i][j] = min(res[i][j], res[i+1][j]+1)
			}
			if j < n-1 {
				res[i][j] = min(res[i][j], res[i][j+1]+1)
			}
		}
	}

	return res
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
