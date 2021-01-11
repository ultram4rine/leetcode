package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isNumber("6e6.5"))
}

func isNumber(s string) bool {
	// remove all spaces from string.
	s = strings.Trim(s, " ")

	// if 's' after trim is empty or doesn't contain numbers return false.
	if s == "" || !strings.ContainsAny(s, "0123456789") {
		return false
	}

	return isNumberRec(s, "0123456789+-.", false, false, false, false)
}

// mustHave means that we need one more symbol, used for 'e'.
func isNumberRec(s string, chars string, wasDigit, wasExp, wasPoint, mustHave bool) bool {
	if mustHave && s == "" {
		return false
	}
	if s == "" {
		if !wasDigit {
			return false
		}

		return true
	}
	if !strings.Contains(chars, s[:1]) {
		return false
	}

	switch string(s[0]) {
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
		if wasExp {
			return isNumberRec(s[1:], "0123456789e", wasDigit, wasExp, true, false)
		}
		return isNumberRec(s[1:], "0123456789e.", true, wasExp, wasPoint, false)
	case "e":
		if wasExp || !wasDigit {
			return false
		}
		return isNumberRec(s[1:], "0123456789+-", wasDigit, true, wasPoint, true)
	case "+", "-":
		return isNumberRec(s[1:], "0123456789.", wasDigit, wasExp, wasPoint, mustHave)
	case ".":
		if wasPoint {
			return false
		}
		return isNumberRec(s[1:], "0123456789e", wasDigit, wasExp, true, false)
	default:
		return false
	}
}
