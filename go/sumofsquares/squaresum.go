package main

import (
	"fmt"
)

func SumofSquare(c, q chan int) {
	y := 1
	for {
		select {
		case c <- (y * y):
			y++
		case <-q:
			return
		}
	}
}

func main() {
	sum := 0
	c := make(chan int)
	q := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			sum += <-c
		}
		fmt.Println(sum)
		q <- 0
	}()
	SumofSquare(c, q)
}
