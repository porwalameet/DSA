package main

import "math"

func mincostTickets(days []int, costs []int) int {
	n := len(days)
	duration := []int{1, 7, 30}
	dp := make([]int, n)
	// for i := range dp {
	//     dp[i] = -1
	// }
	var dfs func(int, int) int
	dfs = func(i, ans int) int {
		if i >= n {
			return 0
		}
		if dp[i] != 0 {
			return dp[i]
		}
		j := i
		for k := 0; k < len(duration); k++ {
			for ; j < n && days[j] < days[i]+duration[k]; j++ {
				//skip
			}
			ans = min(ans, costs[k]+dfs(j, ans))
		}
		dp[i] = ans
		return ans
	}
	return dfs(0, math.MaxInt32)
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	days := []intint{1, 4, 6, 7, 8, 20}
	costs := []int{2, 7, 15}
}
