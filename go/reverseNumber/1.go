package main

import (
	"fmt"
)

func main() {
	a1 := 123
	//logic1
	t := a1
	reversedDigit := 0
	for t > 0 {
		lastDigit := t % 10
		reversedDigit = reversedDigit*10 + lastDigit
		t = t / 10

	}
	fmt.Println(reversedDigit)
}
