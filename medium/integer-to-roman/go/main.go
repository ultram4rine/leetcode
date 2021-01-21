package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("III", intToRoman(3))
	fmt.Println("IV", intToRoman(4))
	fmt.Println("IX", intToRoman(9))
	fmt.Println("LVIII", intToRoman(58))
	fmt.Println("MCMXCIV", intToRoman(1994))
}

func intToRoman(num int) string {
	var (
		ans  string
		rank int
	)
	for num != 0 {
		digit := num % 10
		num /= 10
		s := converter(digit * int(math.Pow10(rank)))
		if s == "" {
			if digit > 5 {
				s += converter(5 * int(math.Pow10(rank)))
				for i := 1; i <= digit-5; i++ {
					s += converter(1 * int(math.Pow10(rank)))
				}
			} else {
				for i := 1; i <= digit; i++ {
					s += converter(1 * int(math.Pow10(rank)))
				}
			}
		}
		s += ans
		ans = s
		rank++
	}
	return ans
}

func converter(n int) string {
	switch n {
	case 1:
		return "I"
	case 4:
		return "IV"
	case 5:
		return "V"
	case 9:
		return "IX"
	case 10:
		return "X"
	case 40:
		return "XL"
	case 50:
		return "L"
	case 90:
		return "XC"
	case 100:
		return "C"
	case 400:
		return "CD"
	case 500:
		return "D"
	case 900:
		return "CM"
	case 1000:
		return "M"
	default:
		return ""
	}
}
