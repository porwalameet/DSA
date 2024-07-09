package main

import "fmt"

// var sum = 0

func main() {

	fmt.Println(addDigit(123565))
}

func addDigit(num int) int {
	if num < 10 {
		return num
	}
	for num > 10 {
		num = sum(num)
	}
	return num

}
func sum(num int) int {

	output := 0
	for num > 0 {
		t := num % 10
		output += t
		num = num / 10
	}
	return output
}
