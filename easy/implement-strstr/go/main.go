package main

import "fmt"

func main() {
	fmt.Println(strStr("mississippi", "sippj"))
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	var index = -1

	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			index = i
			i++
			j := 1
			if j == len(needle) {
				return index
			}
			if i == len(haystack) {
				return -1
			}
			for haystack[i] == needle[j] {
				if j+1 == len(needle) {
					return index
				}
				if i+1 == len(haystack) {
					return -1
				}
				i++
				j++
			}
			i = index
			index = -1
		}
	}

	return index
}
