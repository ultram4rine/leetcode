package main

import (
	"fmt"
	"math"
)

func main() {
	s := "21474836460"
	fmt.Println(myAtoi(s))
}

func myAtoi(s string) int {
	var (
		// answer.
		x int

		// sign.
		negative bool

		// if digit found set it true.
		onlyDigits bool
	)

	// range over given string.
	for _, c := range s {
		if onlyDigits {
			if isDigit(c) {
				if x > math.MaxInt32 {
					return math.MaxInt32
				}
				if x < math.MinInt32 {
					return math.MinInt32
				}

				if negative {
					x = x*10 - int(c-'0')
				} else {
					x = x*10 + int(c-'0')
				}

				if x > math.MaxInt32 {
					return math.MaxInt32
				}
				if x < math.MinInt32 {
					return math.MinInt32
				}
			} else {
				break
			}
		} else {
			if isSpace(c) {
				continue
			} else if isPlus(c) {
				onlyDigits = true
			} else if isMinus(c) {
				onlyDigits = true
				negative = true
			} else if isDigit(c) {
				// if we not yet in onlyDigits mode then we meet digit at first time.
				onlyDigits = true
				x += int(c - '0')
			} else {
				break
			}
		}
	}

	return x
}

func isSpace(c rune) bool {
	return c == 32
}

func isPlus(c rune) bool {
	return c == 43
}

func isMinus(c rune) bool {
	return c == 45
}

func isDigit(c rune) bool {
	return c == 48 || c == 49 || c == 50 || c == 51 || c == 52 || c == 53 || c == 54 || c == 55 || c == 56 || c == 57
}
