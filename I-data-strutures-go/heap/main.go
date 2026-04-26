package main

import (
	"container/heap"
	"fmt"
)

// ===== Heap =====
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	val := old[len(old)-1]
	*h = old[:len(old)-1]
	return val
}

// ===== Kth Largest =====
func findKthLargest(nums []int, k int) int {
	h := &MinHeap{}
	heap.Init(h)

	for _, v := range nums {
		heap.Push(h, v)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return (*h)[0]
}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}

	fmt.Println("Kth Largest:", findKthLargest(nums, 2))
}
