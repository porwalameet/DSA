package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Node struct {
	Vertex int
	Time   int
}

func networkDelay(times [][]int, n, k int) int {
	distance := make([]int, n+1)
	adj := make(map[int][]Node)
	for i := 0; i < len(times); i++ {
		src := times[i][0]
		dst := times[i][1]
		time := times[i][2]

		adj[src] = append(adj[src], Node{Vertex: dst, Time: time})
	}

	for i := 1; i <= n; i++ {
		distance[i] = math.MaxInt64
	}

	minHeap := &MinHeap{}
	heap.Push(minHeap, Node{Vertex: k, Time: 0})
	distance[k] = 0
	for minHeap.Len() > 0 {
		node := heap.Pop(minHeap).(Node)
		vertex := node.Vertex
		time := node.Time
		if distance[vertex] < time {
			continue
		}
		for i := 0; i < len(adj[vertex]); i++ {
			nextNode := adj[vertex][i]
			nextVertex := nextNode.Vertex
			nextTime := nextNode.Time
			if time+nextTime < distance[nextVertex] {
				distance[nextVertex] = time + nextTime
				heap.Push(minHeap, Node{Vertex: nextVertex, Time: distance[nextVertex]})
			}
		}
	}

	maxTime := 0
	for i := 0; i < len(distance); i++ {
		maxTime = max(maxTime, distance[i])
	}
	if maxTime == math.MaxInt64 {
		return -1
	}
	return maxTime
}

type MinHeap []Node

func (h MinHeap) Len() int {
	return len(h)
}

// Sort on time
func (h MinHeap) Less(i, j int) bool {
	return h[i].Time < h[j].Time
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(a interface{}) {
	*h = append(*h, a.(Node))
}

func (h *MinHeap) Pop() interface{} {
	l := len(*h)
	res := (*h)[l-1]
	*h = (*h)[:l-1]
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	n := 2
	k := 2
	times := [][]int{{1, 2, 1}}
	//times := [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}
	fmt.Println(networkDelay(times, n, k))
}
