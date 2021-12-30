package main

func main() {}

func canConstruct(ransomNote string, magazine string) bool {
	hashTable := make([]int, 26)
	for _, c := range magazine {
		hashTable[c-'a']++
	}
	for _, c := range ransomNote {
		hashTable[c-'a']--
		if hashTable[c-'a'] < 0 {
			return false
		}
	}
	return true
}
