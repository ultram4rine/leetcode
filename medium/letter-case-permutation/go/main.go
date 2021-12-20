package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(letterCasePermutation("a1b1"))
}

func letterCasePermutation(s string) []string {
	var res = make([]string, 0)

	s = strings.ToLower(s)
	dfs(s, &res, 0)

	return res
}

func dfs(s string, res *[]string, i int) {
	if i < len(s) {
		dfs(s, res, i+1)
		if unicode.IsLetter(rune(s[i])) {
			s = s[:i] + string(unicode.ToUpper(rune(s[i]))) + s[i+1:]
			dfs(s, res, i+1)
		}
	} else {
		*res = append(*res, s)
	}
}
