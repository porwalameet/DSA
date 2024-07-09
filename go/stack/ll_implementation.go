package main

type Node struct {
	value int
	next  *Node
}

type Stack struct {
	root   *Node
	Length int
}

func (s *Stack) Push(value interface{}) bool {
	newNode := &Node{
		value: value,
		next:  nil,
	}
	if s.Length == 0 {
		s.root = newNode
		s.Length++
	}

}
