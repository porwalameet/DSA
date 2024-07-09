package main

import (
	"fmt"
)

func getPosition(n []int, key int) []int {

	if len(n) == 0 {
		return []int{-1, -1}
	}
	pos := binarySearch(n, 0, len(n)-1, key)
	fmt.Println(pos)
	if pos == -1 {
		return []int{-1, -1}
	}
	start, end := pos, pos
	var temp1, temp2 int
	for start != -1 {
		temp1 = start
		start = binarySearch(n, 0, start-1, key)
	}
	start = temp1
	for end != -1 {
		temp2 = end
		end = binarySearch(n, end+1, len(n)-1, key)
	}
	end = temp2
	return []int{start, end}
}

func binarySearch(n []int, left, right, key int) int {
	for left <= right {
		mid := (left + right) / 2
		if key == n[mid] {
			return mid
		} else if key > n[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	return -1
}

func main() {
	fmt.Println(getPosition([]int{3, 4, 5, 5, 5, 5, 5, 5, 7, 9, 10, 11}, 10))
}
