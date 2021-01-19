package main

func main() {}

func xorOperation(n int, start int) int {
	var answer int
	for i := 0; i < n; i++ {
		answer ^= start + 2*i
	}
	return answer
}
