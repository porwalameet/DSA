package main

import "fmt"

func quicksort(n []int, left, right, k int) int {
	if left < right {
		partIdx := getPartition(n, left, right)
		if partIdx == k {
			return n[partIdx]
		} else if partIdx > k {
			quicksort(n, left, partIdx-1, k)
		} else {
			quicksort(n, partIdx+1, right, k)
		}
	}
}

func getPartition(n []int, left, right int) int {
	pivotEle := n[right]
	partIdx := left
	for j := left; j < right; j++ {
		if n[j] < pivotEle {
			swap(n, partIdx, j)
			partIdx++
		}

	}
	swap(n, partIdx, right)
	return partIdx
}

func swap(n []int, i, j int) {
	n[i], n[j] = n[j], n[i]
}

func kthlargestEle(n []int, k int) int {
	// index := len(n) - k
	fmt.Println(quicksort(n, 0, len(n)-1, k))
	// return n[index]
}

func main() {
	fmt.Println(kthlargestEle([]int{5, 3, 1, 6, 4, 2}, 4))
}
