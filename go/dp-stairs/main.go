package main

import "fmt"

func minCostN(n []int) {
	// dp := make([]int, len(n)+1)
	// l := len(n)
	// for i := 0; i < l; i++ {
	// 	if i < 2 {
	// 		dp[i] = n[i]
	// 	} else {
	// 		dp[i] = n[i] + min(dp[i-1], dp[i-2])
	// 	}
	// }
	// fmt.Println(min(dp[l-1], dp[l-2]))
	l := len(n)
	dpOne, dpTwo := n[0], n[1]
	for i := 2; i < l; i++ {
		current := n[i] + min(dpOne, dpTwo)
		dpOne = dpTwo
		dpTwo = current
	}

	fmt.Println(min(dpOne, dpTwo))

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	minCostN([]int{15, 10, 25, 30})
}
