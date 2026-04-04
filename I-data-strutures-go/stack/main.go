package main

import (
	"fmt"
	"sync"
)

type Stack struct {
	lock sync.Mutex
	s    []int
}

func newStack() *Stack {
	return &Stack{lock: sync.Mutex{}, s: make([]int, 0)}
}

func (st *Stack) Push(v int) {
	defer st.lock.Unlock()
	st.lock.Lock()
	st.s = append(st.s, v)
}

func (st *Stack) Pop() int {
	defer st.lock.Unlock()
	st.lock.Lock()
	if len(st.s) == 0 {
		return 0
	}
	res := st.s[len(st.s)-1]
	st.s = st.s[:len(st.s)-1]
	return res
}
func main() {
	stack := newStack()
	stack.Push(3)
	stack.Push(4)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
