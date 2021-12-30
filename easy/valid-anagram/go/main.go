package main

func main() {}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hashTable := make([]int, 26)

	for _, c := range s {
		hashTable[c-'a']++
	}

	for _, c := range t {
		hashTable[c-'a']--
	}

	for _, v := range hashTable {
		if v < 0 {
			return false
		}
	}

	return true
}
