package main

import (
	"fmt"
)

//

func missingNum(nums []int) int {
	n := len(nums) + 1
	sum := n * (n + 1) / 2
	for _, num := range nums {
		sum -= num
	}
	return sum
}

func main() {
	fmt.Println(missingNum([]int{1, 5, 2, 3, 6, 7, 4, 8, 10}))
}
