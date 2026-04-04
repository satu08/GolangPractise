package main

import (
	"fmt"
	"sync"
)

func main() {
	s := []int{1, 2}
	s = append(s, 3)
	s = modify(s)
	fmt.Println(s)
	fmt.Println(reverseString("satyanarayan"))
	arr := [3]any{"satya", 3, 24.89}
	fmt.Println(arr)
	fmt.Println(validParanthesis("([])"))
	stack := newStack()
	stack.Push(3)
	stack.Push(4)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
func modify(s []int) []int {
	s = append(s, 100)
	return s
}

func reverseString(s string) string {
	b := []rune(s)
	j := len(s) - 1
	i := 0
	for i <= j {
		temp := b[i]
		b[i] = b[j]
		b[j] = temp
		i++
		j--
	}
	return string(b)
}

func validParanthesis(s string) bool {
	run := []rune(s)
	closingMap := make(map[string]string)

	closingMap[")"] = "("
	closingMap["}"] = "{"
	closingMap["]"] = "["
	var tempStack []string
	for _, ch := range run {
		if _, ok := closingMap[string(ch)]; ok {
			if len(tempStack) > 0 {
				if tempStack[len(tempStack)-1] == closingMap[string(ch)] {
					tempStack = tempStack[:len(tempStack)-1]
				} else {
					return false
				}
			} else {
				return false
			}
		} else {
			tempStack = append(tempStack, string(ch))
		}
	}
	if len(tempStack) != 0 {
		return false
	}
	return true
}

func removeDuplicates(nums []int) int {
	length := len(nums)
	res := 1
	for i := 1; i < length; i++ {
		if nums[i] != nums[res-1] {
			nums[res] = nums[i]
			res++

		}
	}
	return res
}

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
