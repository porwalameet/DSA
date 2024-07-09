package main

import (
	"fmt"
	"strconv"
)

//"121" shd give 121

func main() {
	s := "-451"
	fmt.Println(myAtoi(s))
}
func myAtoi(s string) int {
	positive := true
	if string(s[0]) == "-" {
		positive = false
		s = s[1:]
	}
	output := 0
	for i := range s {
		tempNum, _ := strconv.Atoi(string(s[i]))
		output = output*10 + tempNum
	}
	if !positive {
		output *= -1
	}
	return output
}
