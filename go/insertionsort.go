package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateSlice(size int) []int {
	s := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		s[i] = rand.Intn(9) - rand.Intn(9)
	}
	return s

}
func main() {
	s := generateSlice(5)
	fmt.Println("Before: ", s)
	//insertionSort(s)
	s = mergeSort(s)
	fmt.Println("After: ", s)
}

func mergeSort(s []int) []int {
	length := len(s)
	if length == 1 {
		return s
	}
	middle := int(length / 2)
	var (
		leftSlice  = make([]int, middle)
		rigthSlice = make([]int, length-middle)
	)
	leftSlice = s[0:middle]
	rigthSlice = s[middle:length]
	return merge(mergeSort(leftSlice), mergeSort(rigthSlice))

}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	fmt.Println(result)
	return result
}

func insertionSort(s []int) {
	var n = len(s)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if s[j-1] > s[j] {
				s[j-1], s[j] = s[j], s[j-1]
			}
			j--
		}
	}
}
