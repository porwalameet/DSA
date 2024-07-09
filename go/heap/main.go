package main

import "fmt"

//MaxHeap Struct has a slice that holds array

type MaxHeap struct {
	array []int
}

// Insert adds an element to the heap.
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)

	}
}

func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if h.array[l] > h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}
		//compare and swap
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

func (h *MaxHeap) Extract() int {
	if len(h.array) == 0 {
		fmt.Println("heap size is 0")
		return -1
	}

	extracted := h.array[0]
	length := len(h.array)
	h.array[0] = h.array[length-1]
	h.array = h.array[:length-1]
	h.maxHeapifyDown(0)
	return extracted
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 2
}
func right(i int) int {
	return 2*i + 2
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {
	m := &MaxHeap{}
	fmt.Println(m)
	buildHeap := []int{1, 5, 4, 3, 2, 10, 20, 30}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}
	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)

	}

}
