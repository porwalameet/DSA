package main

import "fmt"

// Example:1
// reverse a slice
func reverseInt(n []int) {
	for i, j := 0, len(n)-1; i < j; {
		n[i], n[j] = n[j], n[i]
		i++
		j--
	}
	fmt.Println(n)
}

// reverse a slice
func reverseString(s string) {
	sRune := []rune(s)
	for i, j := 0, len(sRune)-1; i < j; {
		sRune[i], sRune[j] = sRune[j], sRune[i]
		i++
		j--
	}
	fmt.Println(string(sRune))
}

// Find Max Min from Array

func findMaxMin(n []int) {
	max, min := -99999, 99999
	maxIndex, minIndex := -1, -1
	for i := 0; i < len(n); i++ {
		if n[i] > max {
			max = n[i]
			maxIndex = i
		}
		if n[i] < min {
			min = n[i]
			minIndex = i
		}
	}

	fmt.Printf("Max:%d, MaxIndex: %d, Min: %d, MinIndex: %d", max, maxIndex, min, minIndex)
}

// Quick Select to find kth larget element
func quickSort(n []int, l, r, k int) int {
	for {
		pos := partition(n, l, r)
		if pos == k {
			return n[pos]
		} else if k < pos {
			r = pos - 1
		} else {
			l = pos + 1
		}
	}
}

func swap(n []int, i, j int) {
	n[i], n[j] = n[j], n[i]
}

func partition(n []int, l, r int) int {
	pivot := n[r]
	i := l
	for j := l; j < r; j++ {
		if n[j] <= pivot {
			swap(n, i, j)
			i++
		}
	}
	swap(n, i, r)
	return i
}

func kthlarge(n []int, k int) {
	index := len(n) - k
	// quickSort(n, 0, len(n)-1)
	fmt.Println(quickSort(n, 0, len(n)-1, index))
}

func findEle(mat [][]int, r, c, key int) {
	i, j := 0, c-1
	for i < r && j >= 0 {
		if mat[i][j] == key {
			fmt.Println(i, j)
			break
		} else if mat[i][j] > key {
			j -= 1
		} else {
			i += 1
		}
	}
	if i == r || j < 0 {
		fmt.Println("Not found")
	}
}

func first(row []int, l, r int) int {
	for l <= r {
		mid := l + (r-l)/2
		fmt.Println("mid: ", mid)
		if row[mid] == 1 && (mid == 0 || row[mid-1] == 0) {
			return mid
		} else if row[mid] == 0 {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
func findMax1sRow(mat [][]int, m, n int) {
	rowWithMax1s := 0
	max := -1
	for i := range mat {
		currentIndex := first(mat[i], 0, n-1)
		fmt.Println("current index: ", currentIndex)
		if currentIndex != -1 && n-currentIndex > max {
			rowWithMax1s = i
			max = n - currentIndex
		}
	}
	fmt.Printf("Row with max 1s :%d ", rowWithMax1s)
	fmt.Println()

}

func main() {
	//Example 1
	// reverseInt([]int{4, 3, 5, 2})
	// reverseString("Ameet")
	//Example 2
	// findMaxMin([]int{1000, 11, 445, 1, 330, 3000})

	// kth largest element
	// kthlarge([]int{1000, 11, 445, 1, 330, 3000}, 5)
	//matrix :=make(map[int]int)
	m, n := 4, 4
	// var matrix = [][]int{{1, 11, 12, 13}, {14, 15, 17, 24}, {26, 29, 30, 41}, {44, 51, 52, 63}}
	// findEle(matrix, m, n, 17)
	var mat = [][]int{{0, 0, 0, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {0, 0, 1, 1}}
	findMax1sRow(mat, m, n)
}
