package main

import "fmt"

type pair struct {
	row, col, dist int
}

func minDistance(mat [][]int) int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return -1
	}

	rows, cols := len(mat), len(mat[0])
	queue := make([]pair, 0)
	visited := make(map[int]bool)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if mat[i][j] == 1 {
				queue = append(queue, pair{i, j, 0})
				visited[i*cols+j] = true
			}
		}
	}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		for _, d := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			newR, newC := curr.row+d[0], curr.col+d[1]
			if newR >= 0 && newR < rows && newC >= 0 && newC < cols && !visited[newR*cols+newC] {
				if mat[newR][newC] == 0 {
					queue = append(queue, pair{newR, newC, curr.dist + 1})
					visited[newR*cols+newC] = true
				} else if mat[newR][newC] == 1 {
					return curr.dist
				}
			}
			fmt.Println(queue)
		}
	}
	return -1
}

func main() {
	mat := [][]int{
		{1, 0, 1, 0},
		{0, 0, 0, 1},
	}
	fmt.Println("min distance to connected islands:", minDistance(mat))
}
