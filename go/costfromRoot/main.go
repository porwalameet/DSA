package main

import "fmt"

func costfromRoot(edges [][]int, n int) []int {
	g := make([][]int, n)
	ans := make([]int, n)
	size := make([]int, n)
	for _, e := range edges {
		a, b := e[0], e[1]
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	var dfs1 func(int, int, int)
	dfs1 = func(i, fa, d int) {
		ans[0] += d
		size[i] = 1
		for _, j := range g[i] {
			if j != fa {
				dfs1(j, i, d+1)
				size[i] += size[j]
			}
		}
	}

	var dfs2 func(i, fa, t int)
	dfs2 = func(i, fa, t int) {
		ans[i] = t
		for _, j := range g[i] {
			if j != fa {
				dfs2(j, i, t+n-2*size[j])
			}
		}
	}

	dfs1(0, -1, 0)
	dfs2(0, -1, ans[0])
	return ans

}
func main() {
	fmt.Println(costfromRoot([][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}}, 6))
}
