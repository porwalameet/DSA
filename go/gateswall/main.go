package main

import (
	"fmt"
	"math"
)

const (
	WALL = -1
	GATE = 0
	INF  = math.MaxInt32
)

var directions = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func wallgates(m [][]int) {
	for r := 0; r < len(m); r++ {
		for c := 0; c < len(m[0]); c++ {
			if m[r][c] == GATE {
				dfs(m, r, c, 0)
			}
		}
	}

}

func dfs(m [][]int, r, c, cost int) {
	if r < 0 || r >= len(m) || c < 0 || c >= len(m[0]) {
		return
	}
	if m[r][c] < cost {
		return
	}
	m[r][c] = cost
	for _, dir := range directions {
		dfs(m, r+dir[0], c+dir[1], cost+1)
	}

}
func main() {
	m := [][]int{
		{INF, -1, 0, INF},
		{INF, INF, INF, -1},
		{INF, -1, INF, -1},
		{0, -1, INF, INF},
	}
	wallgates(m)
	for r := 0; r < len(m); r++ {
		for c := 0; c < len(m[0]); c++ {
			fmt.Printf("%d ", m[r][c])
		}
		fmt.Println()
	}
}
