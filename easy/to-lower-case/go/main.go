package main

func main() {}

// let's not use strings.ToLower().
func toLowerCase(str string) string {
	var answer string
	for _, c := range str {
		if c >= 'A' && c <= 'Z' {
			c += 'a' - 'A'
		}
		answer += string(c)
	}
	return answer
}
