package main

import "fmt"

type Tree struct {
	root *Node
}

type Node struct {
	data  int
	left  *Node
	right *Node
}

func (t *Tree) insert(data int) *Tree {
	if t.root == nil {
		t.root = &Node{data: data}
	} else {
		t.root.insert(data)
	}
	return t
}

func (n *Node) insert(data int) {
	if data < n.data {
		if n.left == nil {
			n.left = &Node{data: data}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{data: data}
		} else {
			n.right.insert(data)
		}
	}
}

func (t *Tree) preorder() {
	if t.root == nil {
		return
	} else {
		t.root.preorder()
	}
}

func (n *Node) preorder() {
	fmt.Printf("%d  ", n.data)
	if n.left != nil {
		n.left.preorder()
	}
	if n.right != nil {
		n.right.preorder()
	}
}

func (t *Tree) rightView() []int {
	res := []int{}
	if t.root != nil {
		res = t.root.rightView(0, res)
	}
	return res
}

func (n *Node) rightView(level int, res []int) []int {
	if level >= len(res) {
		res = append(res, n.data)
	}
	if n.right != nil {
		res = n.right.rightView(level+1, res)
	}
	if n.left != nil {
		res = n.left.rightView(level+1, res)
	}
	return res
}

func main() {
	t := Tree{}
	t.insert(20)
	t.insert(30)
	t.insert(25)
	t.insert(10)
	t.insert(12)
	t.insert(8)
	t.insert(6)
	t.insert(5)
	t.insert(22)
	t.insert(23)
	t.preorder()
	fmt.Println("Rigth View")
	fmt.Println(t.rightView())
}
