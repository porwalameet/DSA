package main

import "fmt"

type UnionFind struct {
	parent, rank []int
	max          int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n+1)
	rank := make([]int, n+1)
	for i := range parent {
		parent[i] = i
	}
	return &UnionFind{parent, rank, 1}
}

func (uf *UnionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) union(x, y int) {
	rootX, rootY := uf.find(x), uf.find(y)
	if rootX != rootY {
		if uf.rank[rootX] < uf.rank[rootY] {
			rootX, rootY = rootY, rootX
		}
		uf.parent[rootY] = rootX
		if uf.rank[rootX] == uf.rank[rootY] {
			uf.rank[rootX]++
		}
		count := 0
		for _, rank := range uf.rank {
			if rank > count {
				count = rank
			}
		}
		if count > uf.max {
			uf.max = count
		}
	}
}

func minTime(n int, friendships [][]int) int {
	uf := NewUnionFind(n)
	for _, friendship := range friendships {
		uf.union(friendship[0], friendship[1])
	}
	return uf.max - 1
}

func main() {
	fmt.Println(minTime(4, [][]int{{1, 2}, {2, 3}, {3, 4}})) // Output: 3
}
