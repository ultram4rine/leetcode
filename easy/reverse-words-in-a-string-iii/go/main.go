package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}

func reverseWords(s string) string {
	var (
		ans      []rune
		appendTo int
	)

	for i, c := range s {
		if c == ' ' {
			ans = append(ans, ' ')
			appendTo = i + 1
		} else {
			ans = append(ans[:appendTo], append([]rune{c}, ans[appendTo:]...)...)
		}
	}

	return string(ans)
}
