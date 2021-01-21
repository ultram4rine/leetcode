package main

import "fmt"

func main() {
	fmt.Println(2, brokenCalc(2, 3))
	fmt.Println(2, brokenCalc(5, 8))
	fmt.Println(3, brokenCalc(3, 10))
	fmt.Println(39, brokenCalc(1, 100))
}

func brokenCalc(X int, Y int) int {
	ans := 0
	for Y > X {
		if Y%2 == 0 {
			Y /= 2
		} else {
			Y++
		}
		ans++
	}
	return ans + X - Y
}
