package main

func main() {}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var answer = make([]bool, len(candies))
	var max = 0
	for i := range candies {
		if candies[i] > max {
			max = candies[i]
		}
	}
	for i := range candies {
		if candies[i]+extraCandies >= max {
			answer[i] = true
		}
	}
	return answer
}
