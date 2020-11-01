package heap

import (
	"fmt"
)

// Heap holds a heap array
type Heap struct {
	arr []int
}

// New creates a new instance of heap
func New() *Heap {
	return &Heap{
		arr: make([]int, 0),
	}
}

// Insert adds a new element to heap
func (h *Heap) Insert(data int) {
	h.arr = append(h.arr, data)

	size := len(h.arr)

	i := size - 1

	for i != 0 && h.arr[parent(i)] > h.arr[i] {
		h.arr[parent(i)], h.arr[i] = h.arr[i], h.arr[parent(i)]
		i = parent(i)
	}
}

// DecreaseKey decrease a key from heap
func (h *Heap) DecreaseKey(key, val int) error {
	if key >= len(h.arr) {
		return fmt.Errorf("Key is beyond the length")
	}

	h.arr[key] = val
	i := key
	for i != 0 && h.arr[parent(i)] > h.arr[i] {
		h.arr[parent(i)], h.arr[i] = h.arr[i], h.arr[parent(i)]
		i = parent(i)
	}

	return nil
}

// DeleteKey deletes a key from heap
func (h *Heap) DeleteKey(key int) error {
	if key >= len(h.arr) {
		return fmt.Errorf("Key is beyond the length")
	}

	h.DecreaseKey(key, -1)
	h.ExtractMin()

	return nil
}

// ExtractMin removes the min element from heap
func (h *Heap) ExtractMin() int {
	if len(h.arr) == 0 {
		return 0
	}

	min := h.arr[0]
	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	minHeapify(0, len(h.arr), &h.arr)

	return min
}

// minHeapify creates a min heap
func minHeapify(current int, n int, arr *[]int) {
	right := 2*current + 2
	left := 2*current + 1
	root := current

	if left < n && (*arr)[left] < (*arr)[root] {
		root = left
	}

	if right < n && (*arr)[right] < (*arr)[root] {
		root = right
	}

	if root != current {
		(*arr)[current], (*arr)[root] = (*arr)[root], (*arr)[current]
		minHeapify(root, n, arr)
	}
}

// parent returns the parent index of element
func parent(i int) int {
	return (i - 1) / 2
}
