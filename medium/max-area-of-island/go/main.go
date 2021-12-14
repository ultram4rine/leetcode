package main

import "fmt"

func main() {
	/* fmt.Println(maxAreaOfIsland([][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}})) */
	fmt.Println(maxAreaOfIsland([][]int{{0, 1, 0}, {0, 1, 0}}))
}

func maxAreaOfIsland(grid [][]int) int {
	var (
		m       = len(grid)
		n       = len(grid[m-1])
		used    = make([][]bool, m)
		maxArea = 0
	)

	for i := 0; i < m; i++ {
		used[i] = make([]bool, n)
	}

	for i := range grid {
		for j := range grid[i] {
			if used[i][j] {
				continue
			}
			if grid[i][j] == 1 {
				area := calcAreaOfIsland(1, grid, used, i, j, m, n)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}

func calcAreaOfIsland(area int, grid [][]int, used [][]bool, i int, j int, m int, n int) int {
	used[i][j] = true
	for k := 0; k < 4; k++ {
		switch k {
		case 0:
			if j == 0 || grid[i][j-1] == 0 || used[i][j-1] {
				continue
			}
			area++
			area = calcAreaOfIsland(area, grid, used, i, j-1, m, n)
		case 1:
			if i == m-1 || grid[i+1][j] == 0 || used[i+1][j] {
				continue
			}
			area++
			area = calcAreaOfIsland(area, grid, used, i+1, j, m, n)
		case 2:
			if j == n-1 || grid[i][j+1] == 0 || used[i][j+1] {
				continue
			}
			area++
			area = calcAreaOfIsland(area, grid, used, i, j+1, m, n)
		case 3:
			if i == 0 || grid[i-1][j] == 0 || used[i-1][j] {
				continue
			}
			area++
			area = calcAreaOfIsland(area, grid, used, i-1, j, m, n)
		}
	}

	return area
}
