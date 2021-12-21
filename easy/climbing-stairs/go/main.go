package main

import "math"

func main() {}

func climbStairs(n int) int {
	var phi = (float64(1) + math.Sqrt(5)) / float64(2)
	return int(math.Round((math.Pow(phi, float64(n+1)) - (math.Pow(float64(1)-phi, float64(n+1)))) / (math.Sqrt(5))))
}
