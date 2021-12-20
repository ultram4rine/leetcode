package main

func main() {}

func hammingWeight(num uint32) int {
	for num != 0 {
		return 1 + hammingWeight(num&(num-1))
	}
	return 0
}
