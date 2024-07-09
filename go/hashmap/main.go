package main

import (
	"fmt"
)

// Structure for Hashmap
const ArraySize = 7

type hashmap struct {
	array [ArraySize]*bucket
}

// Insert
func (h *hashmap) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// Delete
func (h *hashmap) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)

}

// Search
func (h *hashmap) Search(key string) {
	index := hash(key)
	if h.array[index].search(key) {
		fmt.Println("Found")
	} else {
		fmt.Println("Not found")
	}

}

// Structure for Bucket
type bucket struct {
	head *bucketNode
}

// bucket node structure
type bucketNode struct {
	key  string
	next *bucketNode
}

// Insert
func (b *bucket) insert(key string) {
	newNode := &bucketNode{key: key}
	newNode.next = b.head
	b.head = newNode

}

// Delete
func (b *bucket) delete(key string) {
	if b.head.key == key {
		b.head = b.head.next
		return
	}
	prev := b.head
	current := b.head
	for current.next != nil {
		if current.key == key {
			prev.next = current.next
		}
		prev = current
		current = current.next
	}

}

// Search
func (b *bucket) search(key string) bool {
	current := b.head
	for current != nil {
		if current.key == key {
			return true
		}
		current = current.next
	}
	return false

}

// hash
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

// Init
func Init() *hashmap {
	result := &hashmap{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result

}

func main() {
	testHashTable := Init()
	list := []string{
		"aa",
		"bb",
		"cc",
		"ad",
		"aabc",
		"cd",
		"da",
		"dd",
	}
	for _, val := range list {
		testHashTable.Insert(val)
	}
	testHashTable.Search("cc")
	testHashTable.Delete("cc")
	testHashTable.Search("cc")

}
