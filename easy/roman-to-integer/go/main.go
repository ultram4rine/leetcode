package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	var answer int

	for i := 0; i < len(s); i++ {
		var symbol = string(s[i])
		if symbol == "I" || symbol == "X" || symbol == "C" {
			if i+1 != len(s) {
				switch symbol {
				case "I":
					if string(s[i+1]) == "V" || string(s[i+1]) == "X" {
						symbol = s[i : i+2]
						i++
					}
				case "X":
					if string(s[i+1]) == "L" || string(s[i+1]) == "C" {
						symbol = s[i : i+2]
						i++
					}
				case "C":
					if string(s[i+1]) == "D" || string(s[i+1]) == "M" {
						symbol = s[i : i+2]
						i++
					}
				}
			}
		}
		answer += converter(symbol)
	}

	return answer
}

func converter(s string) int {
	switch string(s) {
	case "I":
		return 1
	case "IV":
		return 4
	case "V":
		return 5
	case "IX":
		return 9
	case "X":
		return 10
	case "XL":
		return 40
	case "L":
		return 50
	case "XC":
		return 90
	case "C":
		return 100
	case "CD":
		return 400
	case "D":
		return 500
	case "CM":
		return 900
	case "M":
		return 1000
	default:
		return 0
	}
}
