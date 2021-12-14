package main

import "fmt"

func main() {
	fmt.Println(floodFill([][]int{{0, 0, 0}, {0, 0, 0}}, 0, 0, 2))
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	var (
		color = image[sr][sc]
		m     = len(image)
		n     = len(image[m-1])
	)

	image[sr][sc] = newColor

	for i := 0; i < 4; i++ {
		switch i {
		case 0:
			if sc == 0 || image[sr][sc-1] != color || image[sr][sc-1] == newColor {
				continue
			}
			floodFill(image, sr, sc-1, newColor)
		case 1:
			if sr == m-1 || image[sr+1][sc] != color || image[sr+1][sc] == newColor {
				continue
			}
			floodFill(image, sr+1, sc, newColor)
		case 2:
			if sc == n-1 || image[sr][sc+1] != color || image[sr][sc+1] == newColor {
				continue
			}
			floodFill(image, sr, sc+1, newColor)
		case 3:
			if sr == 0 || image[sr-1][sc] != color || image[sr-1][sc] == newColor {
				continue
			}
			floodFill(image, sr-1, sc, newColor)
		}
	}

	return image
}
