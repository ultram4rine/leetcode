package main

func main() {}

func combine(n int, k int) [][]int {
	var (
		curr = make([]int, 0)
		last = make([]int, 0)
		res  = make([][]int, 0)
	)

	for j := 1; j <= k; j++ {
		curr = append(curr, j)
		last = append(last, n-k+j)
	}

	var fst = make([]int, len(curr))
	copy(fst, curr)
	res = append(res, fst)

	for {
		if equal(curr, last) {
			break
		}

		var next = getNext(n, k, curr)
		res = append(res, next)
		copy(curr, next)
	}

	return res
}

func equal(a []int, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func getNext(n int, k int, prev []int) []int {
	var next = make([]int, len(prev))
	copy(next, prev)

	for i := k - 1; i >= 0; i-- {
		if next[i] != n-k+i+1 {
			next[i] += 1
			for j := 1; j < k-i; j++ {
				next[i+j] = next[i] + j
			}
			break
		}
	}

	return next
}
