package main

import "fmt"

func dfs(grid [][]int) {
	rows := len(grid)
	if rows == 0 {
		return
	}
	col := len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, col)
	}
	dfsUtil(grid, 0, 0, visited)
}

func dfsUtil(grid [][]int, x, y int, visited [][]bool) {
	r, c := len(grid), len(grid[0])
	if x < 0 || x >= r || y < 0 || y >= c || visited[x][y] {
		return
	}
	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	visited[x][y] = true
	fmt.Print(grid[x][y], " ")
	for _, dir := range directions {
		dfsUtil(grid, x+dir[0], y+dir[1], visited)
	}
	// dfsUtil(grid, x-1, y, visited)
	// dfsUtil(grid, x, y+1, visited)
	// dfsUtil(grid, x+1, y, visited)
	// dfsUtil(grid, x, y-1, visited)
}

func main() {
	grid := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	dfs(grid)
}
