package main

import "fmt"

func travellerWays(m, n int) int {
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	var dfs func() int
	f[1][1] = 1
	dfs = func() int {
		for i := 1; i <= m; i++ {
			for j := 1; j <= n; j++ {
				if i == 1 && j == 1 {
					continue
				}
				f[i][j] = f[i-1][j] + f[i][j-1]
			}

		}
		return f[m][n]
	}
	return dfs()

}

func main() {
	fmt.Println(travellerWays(2, 3))
	fmt.Println(travellerWays(3, 3))
	fmt.Println(travellerWays(5, 5))
	fmt.Println(travellerWays(18, 18))
}
