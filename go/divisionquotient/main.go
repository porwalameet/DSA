package main

import "fmt"

func main() {
	fmt.Println(getQuotient(15, 2))

}

func getQuotient(m, n int) int {
	sign := 1
	if m < 0 || n < 0 {
		sign = -1
	}
	if m < 0 && n < 0 {
		sign = 1
	}
	if m < 0 {
		m = m * -1
	}
	if n < 0 {
		n = n * -1
	}
	start := n
	quotient := 0
	for start <= m {
		quotient++
		start += n
	}

	return quotient * sign
}
