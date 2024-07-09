package workerpool

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Worker represetns a worker that process requests
type Worker struct {
	ID         int
	Wg         *sync.WaitGroup
	ReqHandler map[int]RequestHandler
}

// LaunchWorker starts a worker	for processing incoming requests
func (w *Worker) LaunchWorker(in chan Request, stopCh chan struct{}) {
	go func() {
		defer w.Wg.Done()
		for {
			select {
			case msg, open := <-in:
				if !open {
					// channel is closed, stop processing and return
					fmt.Println("Stopping worker", w.ID)
					return
				}
				w.processRequest(msg)
				time.Sleep(1 * time.Microsecond) //small delay to simulate processing time
			case <-stopCh:
				fmt.Println("Stopping worker", w.ID)
				return
			}
		}
	}()
}

func (w *Worker) processRequest(msg Request) {
	fmt.Printf("Worker %d processing request %v\n", w.ID, msg)
	var handler RequestHandler
	var ok bool
	if handler, ok = w.ReqHandler[msg.Type]; !ok {
		fmt.Printf("Worker %d: No handler for request type %d\n", w.ID, msg.Type)
	} else {
		if msg.Timeout == 0 {
			msg.Timeout = time.Duration(10 * time.Millisecond)
		}
		for attempt := 0; attempt <= msg.Retries; attempt++ {
			var err error
			done := make(chan struct{})
			ctx, cancelt := context.WithTimeout(context.Background(), msg.Timeout)
			defer cancelt()

			go func() {
				err = handler(msg.Data)
				close(done)
			}()

			select {
			case <-done:
				if err == nil {
					return
				}
				fmt.Printf("Worker %d: Error processing request %v: %s\n", w.ID, msg, err)
			case <-ctx.Done():
				fmt.Printf("Worker %d: Timeout processing request %v\n", w.ID, msg)
			}
			fmt.Printf("Worker %d: Retrying request %v\n", w.ID, msg)
		}
		fmt.Printf("worker %d: Failed to process requests %v after %d retries\n", w.ID, msg, msg.MaxRetries)
	}
}
