package main

import (
	"fmt"
)

func canPartition(nums []int, k int) bool {
	s := 0
	for _, n := range nums {
		s += n
	}
	if s%2 != 0 {
		return false
	}

	s = s / 2
	dp := make([]bool, s+1)
	dp[0] = true
	for _, n := range nums {
		for j := s; j >= n; j-- {
			dp[j] = dp[j] || dp[j-n]
			// dp[j] = max(dp[j], dp[j-n]+n)
		}
	}
	return dp[s]
}

func max(a, b int) int {
	if a > b {
		return a

	}
	return b
}

func main() {
	fmt.Println(canPartition([]int{4, 3, 2, 3, 5, 2, 1, 2}, 2))
}
