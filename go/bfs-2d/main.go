package main

import "fmt"

func bfs(grid [][]int) {
	rows := len(grid)
	if rows == 0 {
		return
	}
	col := len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, col)
	}
	// create queue for nots in next direct level
	queue := make([][]int, 0)
	// add initial position as 0,0 to queue to start with
	queue = append(queue, []int{0, 0})

	// traverse each indices in queue
	for len(queue) > 0 {
		x, y := queue[0][0], queue[0][1]
		queue = queue[1:]
		if x < 0 || x >= rows || y < 0 || y >= col || visited[x][y] {
			continue
		}
		directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

		visited[x][y] = true
		fmt.Print(grid[x][y], " ")
		for _, dir := range directions {
			queue = append(queue, []int{x + dir[0], y + dir[1]})
		}
	}
}

func main() {
	grid := [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}}
	bfs(grid)
}
