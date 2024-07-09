package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop
func (s *Stack) Pop() int {
	last := len(s.items) - 1
	popVal := s.items[last]
	s.items = s.items[:last]
	return popVal
}

func main() {
	myStack := Stack{}
	myStack.Push(10)
	myStack.Push(20)
	myStack.Push(30)
	fmt.Println(myStack)
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.items)
}
