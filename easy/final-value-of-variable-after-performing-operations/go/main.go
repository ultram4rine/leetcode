package main

func main() {}

func finalValueAfterOperations(operations []string) int {
	var x = 0

	for _, op := range operations {
		if op == "++X" || op == "X++" {
			x += 1
		} else if op == "--X" || op == "X--" {
			x -= 1
		}
	}

	return x
}
