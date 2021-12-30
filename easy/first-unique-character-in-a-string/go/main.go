package main

func main() {}

func firstUniqChar(s string) int {
	var arr = [26]int{}

	for _, c := range s {
		i := c - 'a'
		arr[i]++
	}

	for i := range s {
		idx := s[i] - 'a'
		if arr[idx] == 1 {
			return i
		}
	}

	return -1
}
