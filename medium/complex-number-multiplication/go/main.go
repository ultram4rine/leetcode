package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(complexNumberMultiply("1+-1i", "1+-1i"))
}

func complexNumberMultiply(num1 string, num2 string) string {
	var (
		num1arr = strings.Split(num1, "+")
		num2arr = strings.Split(num2, "+")

		re1str = num1arr[0]
		im1str = strings.TrimSuffix(num1arr[1], "i")
		re2str = num2arr[0]
		im2str = strings.TrimSuffix(num2arr[1], "i")
	)

	re1, _ := strconv.Atoi(re1str)
	im1, _ := strconv.Atoi(im1str)
	re2, _ := strconv.Atoi(re2str)
	im2, _ := strconv.Atoi(im2str)

	resRe := re1*re2 - im1*im2
	resIm := im1*re2 + re1*im2

	return fmt.Sprintf("%d+%di", resRe, resIm)
}
