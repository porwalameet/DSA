package main

import (
	"fmt"
	"time"
)

func rec(x int) int {
	if x <= 1 {
		return 1
	}
	return x * rec(x-1)
}

func tailrec(x int, total int) int {
	if x == 0 {
		return total
	} else {
		return tailrec(x-1, total*x)
	}
}

func main() {
	start := time.Now().Nanosecond()
	fmt.Println(rec(9))
	fmt.Println(time.Now().Nanosecond() - start)
	start = time.Now().Nanosecond()
	fmt.Println(tailrec(9, 1))
	fmt.Println(time.Now().Nanosecond() - start)

}
