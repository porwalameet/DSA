package main

import "fmt"

func countIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	count := 0
	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 1 {
				count++
				queue := make([][]int, 0)
				queue = append(queue, []int{r, c})
				grid[r][c] = 0
				for len(queue) > 0 {
					cr := queue[0][0]
					cc := queue[0][1]
					queue = queue[1:]
					for _, dir := range directions {
						nextRow := cr + dir[0]
						nextCol := cc + dir[1]
						if nextRow < 0 || nextRow >= len(grid) || nextCol < 0 || nextCol >= len(grid[0]) {
							continue
						}
						if grid[nextRow][nextCol] == 1 {
							queue = append(queue, []int{nextRow, nextCol})
							grid[nextRow][nextCol] = 0
						}
					}
				}
			}
		}
	}
	return count
}

func main() {
	grid := [][]int{{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	}
	fmt.Println(countIsland(grid))

	grid = [][]int{{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
	}
	fmt.Println(countIsland(grid))
}
