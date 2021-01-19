package main

func main() {}

func maximumWealth(accounts [][]int) int {
	maxWealth := -1
	for i := range accounts {
		wealth := 0
		for j := range accounts[i] {
			wealth += accounts[i][j]
		}
		if wealth > maxWealth {
			maxWealth = wealth
		}
	}
	return maxWealth
}
