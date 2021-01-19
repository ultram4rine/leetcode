package main

func main() {}

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	var (
		w1, w2 string
	)
	for i := range word1 {
		w1 += word1[i]
	}
	for i := range word2 {
		w2 += word2[i]
	}
	return w1 == w2
}
