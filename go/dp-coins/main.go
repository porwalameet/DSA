package main

import (
	"fmt"
)

// Unique way to get change based on input coins
// func coinChange(coins []int, amount int) int {
// 	return coinOption(coins, len(coins), amount)
// }

// func coinOption(coins []int, n, amount int) int {
	
// 	if amount == 0 {
// 		return 1
// 	}
// 	if amount < 0 || n <= 0 {
// 		return 0
// 	}
// 	return coinOption(coins, n-1, amount) + coinOption(coins, n, amount-coins[n-1])

// }

// find min coins to sum
func minCoins(coins []int, n, amount int, dp []int) int {
	if amount == 0 {
		return 0
	}
	if dp[amount] != 0 {
		return dp[amount]
	}
	res := 9999
	for i := 0; i < n; i++ {
		if coins[i] <= amount {
			currentMin := minCoins(coins, n, amount-coins[i], dp)
			res = min(res, currentMin+1)
		}
	}
	dp[amount] = res
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// fmt.Println(coinChange([]int{2, 5, 3, 6}, 10))
	amount := 600
	dp := make([]int, amount+1)
	fmt.Println(minCoins([]int{2, 5, 3, 6}, 4, amount, dp))

}
