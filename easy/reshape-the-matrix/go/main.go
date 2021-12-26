package main

func main() {}

func matrixReshape(mat [][]int, r int, c int) [][]int {
	var (
		m = len(mat)
		n = len(mat[0])
	)

	if m*n != r*c || (m == r && n == c) {
		return mat
	}

	var (
		res = make([][]int, r)
		k   = 0
		l   = 0
	)
	for i := range res {
		res[i] = make([]int, c)
		for j := range res[i] {
			res[i][j] = mat[k][l]
			l++
			if l >= n {
				l = 0
				k++
			}
		}
		if k > m {
			k = 0
		}
	}

	return res
}
