package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Pop
func (q *Queue) Dequeue() int {
	dequeueVal := q.items[0]

	q.items = q.items[1:]
	return dequeueVal
}

func main() {
	myQueue := Queue{}
	myQueue.Enqueue(10)
	myQueue.Enqueue(20)
	myQueue.Enqueue(30)
	fmt.Println(myQueue)
	fmt.Println(myQueue.Dequeue())
	fmt.Println(myQueue.items)
}
