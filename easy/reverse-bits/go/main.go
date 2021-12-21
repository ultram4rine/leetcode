package main

func main() {}

func reverseBits(num uint32) uint32 {
	if num == 0 {
		return 0
	}
	var rev uint32 = 0

	for i := 0; i < 32; i++ {
		rev <<= 1
		rev |= num & 1
		num >>= 1
	}

	return rev
}
