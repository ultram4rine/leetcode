package main

func main() {}

func myPow(x float64, n int) float64 {
	// if n is 0 return 1
	if n == 0 {
		return 1.0
	}
	// if n is 1 return x
	if n == 1 {
		return x
	}
	// if x is 0 return 0.
	if x == 0.0 {
		return 0.0
	}
	// if x is 1 return 1.
	if x == 1.0 {
		return 1.0
	}

	var (
		// answer.
		result = 1.0

		// if n negative set it true.
		neg bool
	)

	// let's use exponentiation by squaring.
	if n < 0 {
		n *= -1
		neg = true
	}

	for n > 0 {
		if n%2 == 1 {
			result *= x
		}
		x *= x
		n /= 2
	}

	if neg {
		result = 1 / result
	}

	return result
}
