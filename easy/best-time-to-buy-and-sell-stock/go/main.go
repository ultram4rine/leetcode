package main

func main() {}

func maxProfit(prices []int) int {
	var (
		minPrice  = prices[0]
		maxProfit = 0
	)

	for _, p := range prices {
		if p < minPrice {
			minPrice = p
		} else if p-minPrice > maxProfit {
			maxProfit = p - minPrice
		}
	}

	return maxProfit
}
