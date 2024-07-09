package main

import "fmt"

func romanToInt(s string) int {
	var roman = map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	res := 0
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && roman[string(s[i])] < roman[string(s[i+1])] {
			res = res - roman[string(s[i])]
		} else {
			res = res + roman[string(s[i])]
		}
	}
	return res
}
func main() {
	fmt.Println(romanToInt("XXXIX"))
}
