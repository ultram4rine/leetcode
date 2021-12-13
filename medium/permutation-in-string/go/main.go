package main

import (
	"fmt"
)

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo"))
	fmt.Println(checkInclusion("vnrekkn", "vnrekknjkvervnkjvner"))
	fmt.Println(checkInclusion("hello", "ooolleooleh"))
}

func checkInclusion(s1 string, s2 string) bool {
	var (
		winSize   = len(s1)
		checkHash [26]int
		testHash  [26]int
	)

	for _, c := range s1 {
		checkHash[c-'a'] += 1
	}

outer:
	for i := 0; i+winSize <= len(s2); i++ {
		substr := s2[i : i+winSize]

		for _, c := range substr {
			testHash[c-'a'] += 1
		}

		if checkHash != testHash {
			var newTest [26]int
			testHash = newTest
			continue outer
		}

		return true
	}

	return false
}
