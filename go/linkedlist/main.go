package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type linkedList struct {
	head   *Node
	length int
}

func (l *linkedList) prepend(data int) {
	node := &Node{data: data}
	current := l.head
	l.head = node
	l.head.next = current
	l.length++
}

func (l linkedList) traverse() {
	if l.length == 0 {
		fmt.Println("No elements in linked list")
		return
	}
	current := l.head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}
func (l *linkedList) remove(data int) {
	if l.length == 0 {
		fmt.Println("List is empty")
		return
	}
	if l.head.data == data {
		l.head = l.head.next
		l.length--
		return
	}
	current := l.head
	prev := l.head
	for current != nil {
		if current.data == data {
			prev.next = current.next
			l.length--

		}
		prev = current
		current = current.next
	}
	if current == nil {
		fmt.Println("Value does not exist in list")
	}

}

func main() {
	myList := linkedList{}
	myList.prepend(10)
	myList.prepend(20)
	myList.prepend(30)
	myList.prepend(30)
	myList.traverse()
	fmt.Println("After deletion")
	myList.remove(30)
	myList.remove(300)
	myList.traverse()
}
