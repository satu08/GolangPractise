package main

import (
	"fmt"
	"sync"
)

func modify(s []int) {
	s[0] = 100
}
func appendSlice(s []int) {
	s = append(s, 99) // memory is reallocated
}
func main() {
	s := []int{1, 2, 3}
	modify(s)
	fmt.Println(s) // [100 2 3]
	appendSlice(s)
	fmt.Println(s)

	s1 := make([]int, 0, 4)
	s1 = append(s1, 1, 2)

	a := s1[:1]
	b := s1[1:2]

	a = append(a, 100)
	fmt.Println(b) // 😱 unexpected // both share the same slice
	m := map[string]int{}
	modifyMap(m)
	fmt.Println(m) // map[a:10]

	// sync map used when unkwon keys are there
	var sm sync.Map
	sm.Store("a", 1)
	v, _ := sm.Load("a")
	fmt.Println(v)

	var x *int = nil
	var p interface{}
	p = x
	var p2 interface{}
	p2 = x
	fmt.Println(p, p2)

	var y *int
	var q interface{}
	q = 2
	y = new(int)
	*y = 2
	q = y
	fmt.Println(y == q)

	// Dynamic Array

	arr := NewDynamicArray(2)

	arr.Append(10)
	arr.Append(20)
	arr.Append(30) // triggers resize

	fmt.Println("Len:", arr.Len()) // 3
	fmt.Println("Cap:", arr.Cap()) // 4
	fmt.Println(arr.Get(1))        // 20

}

func modifyMap(m map[string]int) {

	m["a"] = 10
}

// make map concurrent safe using mutex
type SafeMap struct {
	mu sync.RWMutex
	m  map[string]int
}

func (s *SafeMap) Get(k string) int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.m[k]
}

func (s *SafeMap) Set(k string, v int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[k] = v
}

// Dynamic Array example

type DynamicArray struct {
	data []int
	size int
}

func NewDynamicArray(capacity int) *DynamicArray {
	return &DynamicArray{
		data: make([]int, capacity),
		size: 0,
	}
}

func (da *DynamicArray) Append(value int) {
	if da.size == len(da.data) {
		da.resize()
	}
	da.data[da.size] = value
	da.size++
}

func (da *DynamicArray) resize() {
	newCap := 1
	if len(da.data) > 0 {
		newCap = len(da.data) * 2
	}
	newData := make([]int, newCap)
	for i := 0; i < da.size; i++ {
		newData[i] = da.data[i]
	}
	da.data = newData
}

func (da *DynamicArray) Get(index int) int {
	if index < 0 || index >= da.size {
		panic("index out of bounds")
	}
	return da.data[index]
}

func (da *DynamicArray) Len() int {
	return da.size
}

func (da *DynamicArray) Cap() int {
	return len(da.data)
}
