package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	oddCh, evenCh := make(chan bool, 1), make(chan bool, 1)
	wg = sync.WaitGroup{}
	wg.Add(2)

	go even(evenCh, oddCh)
	go odd(oddCh, evenCh)

	// initiate the odd function to start the sequence
	evenCh <- true
	wg.Wait()
}

func even(evenCh, oddCh chan bool) {
	for i := 2; i <= 5; i += 2 {
		<-oddCh
		fmt.Println(i)
		evenCh <- true
	}
	wg.Done()
}

func odd(oddCh, evenCh chan bool) {
	for i := 1; i <= 5; i += 2 {
		<-evenCh
		fmt.Println(i)
		oddCh <- true
	}
	wg.Done()
}
