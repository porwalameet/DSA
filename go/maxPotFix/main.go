package main

import (
	"fmt"
	"sort"
)

func solve(n int, s string, b int) int {
	var continuousPotCount []int
	cur := 0
	for _, c := range s {
		if c == 'x' || c == 'X' {
			cur++
		} else {
			if cur > 0 {
				continuousPotCount = append(continuousPotCount, cur)
				cur = 0
			}
		}
	}
	if cur > 0 {
		continuousPotCount = append(continuousPotCount, cur)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(continuousPotCount)))

	ans := 0
	for _, p := range continuousPotCount {
		if p+1 <= b {
			ans += p
			b -= (p + 1)
		} else {
			ans += max(0, b-1)
			break
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(solve(17, "...xxx..X....xxx.", 7)) // ans = 5
	fmt.Println(solve(7, "..xxxxx", 4))            // ans = 3
	fmt.Println(solve(11, "x.x.xxx...x", 14))      // ans = 6
	fmt.Println(solve(2, "..", 5))                 // ans = 0
}
