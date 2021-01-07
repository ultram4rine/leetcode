package main

import "fmt"

func main() {
	fmt.Println(defangIPaddr("1.1.1.1"))
}

// let's not use strings.ReplaceAll.
func defangIPaddr(address string) string {
	arr := []byte(address)
	for i := 0; i < len(arr); i++ {
		// if dot found.
		if arr[i] == 46 {
			// insert '[' before dot.
			arr = append(arr, 0)
			copy(arr[i:], arr[i-1:])
			arr[i] = 91
			i += 2

			// insert ']' after dot.
			arr = append(arr, 0)
			copy(arr[i:], arr[i-1:])
			arr[i] = 93
			i++
		}
	}
	return string(arr)
}
