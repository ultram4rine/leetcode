package main

import "fmt"

func main() {
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}))
}

func orangesRotting(grid [][]int) int {
	var (
		q    = make([][2]int, 0)
		temp [2]int
		ans  int

		m = len(grid)
		n = len(grid[0])
	)

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 2 {
				q = append(q, [2]int{i, j})
			}
		}
	}

	q = append(q, [2]int{-1, -1})

	for len(q) != 0 {
		var flag = false
		for !isDelim(q[0]) {
			temp = q[0]

			if isValid(temp[0]+1, temp[1], m, n) && grid[temp[0]+1][temp[1]] == 1 {
				if !flag {
					ans++
					flag = true
				}
				grid[temp[0]+1][temp[1]] = 2

				temp[0]++
				q = append(q, temp)
				temp[0]--
			}

			if isValid(temp[0]-1, temp[1], m, n) && grid[temp[0]-1][temp[1]] == 1 {
				if !flag {
					ans++
					flag = true
				}
				grid[temp[0]-1][temp[1]] = 2

				temp[0]--
				q = append(q, temp)
				temp[0]++
			}

			if isValid(temp[0], temp[1]+1, m, n) && grid[temp[0]][temp[1]+1] == 1 {
				if !flag {
					ans++
					flag = true
				}
				grid[temp[0]][temp[1]+1] = 2

				temp[1]++
				q = append(q, temp)
				temp[1]--
			}

			if isValid(temp[0], temp[1]-1, m, n) && grid[temp[0]][temp[1]-1] == 1 {
				if !flag {
					ans++
					flag = true
				}
				grid[temp[0]][temp[1]-1] = 2

				temp[1]--
				q = append(q, temp)
				temp[1]++
			}

			q = q[1:]
		}

		q = q[1:]

		if len(q) != 0 {
			temp[0] = -1
			temp[1] = -1
			q = append(q, temp)
		}
	}

	if checkAll(grid) {
		return -1
	}

	return ans
}

func isValid(i, j, m, n int) bool {
	return i >= 0 && j >= 0 && i < m && j < n
}

func isDelim(temp [2]int) bool {
	return temp[0] == -1 && temp[1] == -1
}

func checkAll(grid [][]int) bool {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				return true
			}
		}
	}
	return false
}
