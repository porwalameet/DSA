package main

import "fmt"

func dfs(currentID int, adjList map[int][]int, informTime []int) int {
	if len(adjList[currentID]) == 0 {
		return 0
	}
	maxTime := 0
	sub := adjList[currentID]
	for e := 0; e < len(sub); e++ {
		maxTime = max(maxTime, dfs(sub[e], adjList, informTime))
	}
	return maxTime + informTime[currentID]

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	n := 8
	headId := 4
	manager := []int{2, 2, 4, 6, -1, 4, 4, 5}
	informTime := []int{0, 0, 4, 0, 7, 3, 6, 0}
	adjList := make(map[int][]int, n)
	for e := 0; e < n; e++ {
		mgr := manager[e]
		if mgr == -1 {
			continue
		}
		adjList[mgr] = append(adjList[mgr], e)

	}
	fmt.Println(dfs(headId, adjList, informTime))
}
