package main

import "fmt"

func findSubArraySum(nums []int) int {
	size := len(nums)
	max := -99999
	current_max := 0
	start, end, s := 0, 0, 0
	for i := 0; i < size; i++ {
		current_max = current_max + nums[i]
		if current_max > max {
			max = current_max
			start = s
			end = i
		}
		if current_max < 0 {
			current_max = 0
			s = i + 1
		}
	}
	for i := start; i <= end; i++ {
		fmt.Printf("%d ", nums[i])
	}
	fmt.Println()
	return max

}

func main() {
	fmt.Println("Max Sum: ", findSubArraySum([]int{1, 2, 3, 4, 5, 6, 7}))
}
