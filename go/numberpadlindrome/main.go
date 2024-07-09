package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindone(12321))
}
func isPalindone(num int) bool {
	s := strconv.Itoa(num)
	for i := 0; i < (len(s) / 2); i++ {
		if string(s[i]) != string(s[len(s)-1-i]) {
			return false
		}
	}
	return true
}

func solution2(num int) bool {
	reverse := 0
	temp := num
	for temp > 0 {
		lastDigit := temp % 10
		reverse = reverse*10 + lastDigit
		temp = temp / 10
	}
	return reverse == num
}
