package main

import "fmt"

func main() {
	fmt.Println(add1([]int{4, 9, 9}))
}
func add1(nums []int) []int {
	output := make([]int, 0)
	length := len(nums)
	lastDigit := nums[length-1]
	add := lastDigit + 1
	carry := add / 10
	lastDigit = add % 10
	output = append(output, lastDigit)
	for i := length - 2; i >= 0; i-- {
		o := nums[i] + carry
		lastDigit = o % 10
		carry = o / 10
		output = append(output, lastDigit)
	}
	if carry == 1 {
		output = append(output, 1)
	}

	return reverse(output, len(output))
}

func reverse(input []int, len int) []int {
	start, end := 0, len-1
	for start < end {
		input[start], input[end] = input[end], input[start]
		start++
		end--
	}
	return input
}
