package main

func main() {}

func decode(encoded []int, first int) []int {
	var answer = make([]int, len(encoded)+1)
	answer[0] = first
	for i := range encoded {
		answer[i+1] = encoded[i] ^ answer[i]
	}
	return answer
}
