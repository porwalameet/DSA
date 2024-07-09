package main

import "fmt"

func lcs(t1, t2 string) int {
	m, n := len(t1), len(t2)
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if t1[i-1] == t2[j-1] {
				f[i][j] = 1 + f[i-1][j-1]
			} else {
				f[i][j] = max(f[i-1][j], f[i][j-1])
			}
		}
	}
	fmt.Println(f)
	return f[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	fmt.Println(lcs("abcd", "ace"))
}
