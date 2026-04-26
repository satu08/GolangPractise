package main

import "fmt"

func main() {

	// 🔹 Deque operations
	d := Deque{}
	d.PushBack(10)
	d.PushBack(20)
	d.PushFront(5)

	fmt.Println("Front:", d.Front())
	fmt.Println("Back:", d.Back())

	d.PopFront()
	fmt.Println("After PopFront:", d.data)

	// 🔹 Sliding Window
	arr := []int{1, 3, -1, -3, 5, 3, 6, 7}
	fmt.Println("Sliding Window:", maxSlidingWindow(arr, 3))
}

// ===== Deque =====
type Deque struct {
	data []int
}

func (d *Deque) PushFront(val int) {
	d.data = append([]int{val}, d.data...)
}

func (d *Deque) PushBack(val int) {
	d.data = append(d.data, val)
}

func (d *Deque) PopFront() int {
	if len(d.data) == 0 {
		return -1
	}
	val := d.data[0]
	d.data = d.data[1:]
	return val
}

func (d *Deque) Front() int {
	if len(d.data) == 0 {
		return -1
	}
	return d.data[0]
}

func (d *Deque) Back() int {
	if len(d.data) == 0 {
		return -1
	}
	return d.data[len(d.data)-1]
}

// ===== Sliding Window =====
func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	deque := []int{}

	for i := 0; i < len(nums); i++ {

		if len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}

		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, i)

		if i >= k-1 {
			res = append(res, nums[deque[0]])
		}
	}
	return res
}
