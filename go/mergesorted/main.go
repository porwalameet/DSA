package main

import "fmt"

func mergeSorted(a, b []int) {
	m, n := len(a), len(b)
	res := make([]int, m+n)
	for n != 0 {
		if m != 0 && a[m-1] > b[n-1] {
			res[m+n-1] = a[m-1]
			m--
		} else {
			res[m+n-1] = b[n-1]
			n--
		}
	}

	fmt.Println(res)
}

func main() {
	// a := make([]int, 10)
	// b := make([]int, 10)
	mergeSorted([]int{1, 2, 5, 6, 7}, []int{1, 5, 6, 7, 9})
}
