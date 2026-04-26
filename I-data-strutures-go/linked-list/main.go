package main

import "fmt"

// Node definition
type ListNode struct {
	Val  int
	Next *ListNode
}

// 🔹 Insert at end
func insert(head *ListNode, val int) *ListNode {
	newNode := &ListNode{Val: val}

	if head == nil {
		return newNode
	}

	curr := head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = newNode

	return head
}

// 🔹 Print list
func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " -> ")
		head = head.Next
	}
	fmt.Println("nil")
}

// 🔹 Reverse linked list
func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// 🔹 Find middle (fast & slow)
func findMiddle(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 🔹 Detect cycle
func hasCycle(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}
	return false
}

// 🔹 Delete a value
func deleteVal(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	curr := dummy

	for curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return dummy.Next
}

// 🔹 MAIN
func main() {
	var head *ListNode

	// Insert
	head = insert(head, 1)
	head = insert(head, 2)
	head = insert(head, 3)
	head = insert(head, 4)

	fmt.Println("Original List:")
	printList(head)

	// Reverse
	head = reverse(head)
	fmt.Println("Reversed:")
	printList(head)

	// Middle
	mid := findMiddle(head)
	fmt.Println("Middle:", mid.Val)

	// Delete
	head = deleteVal(head, 3)
	fmt.Println("After Delete 3:")
	printList(head)

	// Cycle test
	head.Next.Next.Next = head // create cycle
	fmt.Println("Has Cycle:", hasCycle(head))
}
