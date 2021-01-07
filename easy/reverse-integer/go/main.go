package main

import "math"

func main() {}

func reverse(x int) int {
	// if x is zero return zero.
	if x == 0 {
		return 0
	}

	var (
		// answer.
		y int

		// length of x.
		l = getLength(x)
	)

	// get digits of x and pass it to y.
	for x != 0 {
		digit := x % 10
		x /= 10

		y += digit * int(math.Pow10(l))
		l--
	}

	// if x negative make y negative too.
	if x < 0 {
		y *= -1
	}

	// if y out of int32 range return 0.
	if y > int(math.Pow(2, 31))-1 || y < -int(math.Pow(2, 31)) {
		return 0
	}

	return y
}

// Log based length getting.
func getLength(x int) int {
	// if x is negative make it positive.
	if x < 0 {
		x *= -1
	}

	return int(math.Floor(math.Log10(float64(x) + 1)))
}
