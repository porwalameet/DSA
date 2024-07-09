package main

import (
	"fmt"
)

func find(parent []int, x int) int {
	if parent[x] != x {
		parent[x] = find(parent, parent[x])
	}
	return parent[x]
}

func union(parent, rank []int, x, y int) {
	px := find(parent, x)
	py := find(parent, y)
	if px != py {
		if rank[px] > rank[py] {
			parent[py] = px
		} else if rank[px] < rank[py] {
			parent[px] = py
		} else {
			parent[px] = py
			rank[py]++
		}
	}
}

func earliestTime(logs [][]int, N int) int {
	parent := make([]int, N)
	rank := make([]int, N)
	for i := range parent {
		parent[i] = i
	}

	for _, log := range logs {
		ts, x, y := log[0], log[1], log[2]
		union(parent, rank, x, y)

		// check if members are connected
		root := find(parent, 0)
		allConnected := true
		for i := 1; i < N; i++ {
			if find(parent, i) != root {
				allConnected = false
				break
			}
		}
		if allConnected {
			fmt.Println(parent, rank)
			return ts
		}
	}
	return -1

}

func main() {
	logs := [][]int{{1, 0, 1},
		{2, 1, 2},
		{3, 0, 2},
		{4, 1, 3},
		{5, 2, 4},
		{6, 3, 4},
		{7, 5, 6},
		{8, 5, 4}}
	N := 7
	fmt.Println(earliestTime(logs, N))
}
