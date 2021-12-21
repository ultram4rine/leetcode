package main

func main() {}

func minimumTotal(triangle [][]int) int {
	var n = len(triangle)
	if n == 1 {
		return triangle[0][0]
	}

	var memo = make([]int, len(triangle[n-1]))
	for i := 0; i < len(triangle[n-1]); i++ {
		memo[i] = triangle[n-1][i]
	}

	for i := n - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			memo[j] = triangle[i][j] + minInt(memo[j], memo[j+1])
		}
	}

	return memo[0]
}

func minInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
