package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	top  int
	list []interface{}
}

func (stack *Stack) Push(value interface{}) {
	stack.list = append(stack.list, value)
	stack.top++
}

func (stack *Stack) Pop() (value interface{}, err error) {
	if stack.top <= 0 {
		return nil, errors.New("Stack is empty")
	}
	value = stack.list[stack.top-1]
	stack.list = stack.list[:stack.top-1]
	stack.top--
	return value, nil
}

func (stack *Stack) Length() (value int) {
	return stack.top
}

func (stack *Stack) Display() (list []interface{}) {
	return stack.list
}
func main() {
	s := Stack{}
	s.Push(10)
	s.Push(2)
	fmt.Println(s.Length())
	fmt.Println(s.Display())
	s.Pop()
	fmt.Println(s.Display())
}
