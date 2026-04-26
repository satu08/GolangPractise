package main

import "fmt"

func main() {

	// 🔹 Stack usage
	s := Stack{}
	s.Push(10)
	s.Push(20)
	s.Push(30)

	fmt.Println("Pop:", s.Pop())
	fmt.Println("Peek:", s.Peek())

	// 🔹 Valid Parentheses
	fmt.Println("Valid:", isValid("(){}[]"))

	// 🔹 Next Greater
	arr := []int{1, 3, 2, 4}
	fmt.Println("Next Greater:", nextGreater(arr))
}

// ===== Stack =====
type Stack struct {
	data []int
}

func (s *Stack) Push(val int) {
	s.data = append(s.data, val)
}

func (s *Stack) Pop() int {
	if len(s.data) == 0 {
		return -1
	}
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val
}

func (s *Stack) Peek() int {
	if len(s.data) == 0 {
		return -1
	}
	return s.data[len(s.data)-1]
}

// ===== Valid Parentheses =====
func isValid(s string) bool {
	stack := []rune{}

	for _, ch := range s {
		if ch == '(' || ch == '{' || ch == '[' {
			stack = append(stack, ch)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if (ch == ')' && top != '(') ||
				(ch == '}' && top != '{') ||
				(ch == ']' && top != '[') {
				return false
			}
		}
	}
	return len(stack) == 0
}

// ===== Next Greater =====
func nextGreater(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	stack := []int{}

	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] <= nums[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			res[i] = -1
		} else {
			res[i] = stack[len(stack)-1]
		}

		stack = append(stack, nums[i])
	}
	return res
}
