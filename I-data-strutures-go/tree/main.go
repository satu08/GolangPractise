package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	// Build tree
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	fmt.Print("Preorder: ")
	preorder(root)

	fmt.Print("\nInorder: ")
	inorder(root)

	fmt.Print("\nPostorder: ")
	postorder(root)

	fmt.Print("\nLevel Order: ")
	levelOrder(root)

	fmt.Println("\nHeight:", height(root))
}

// ===== Traversals =====
func preorder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Val, " ")
	preorder(root.Left)
	preorder(root.Right)
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	fmt.Print(root.Val, " ")
	inorder(root.Right)
}

func postorder(root *TreeNode) {
	if root == nil {
		return
	}
	postorder(root.Left)
	postorder(root.Right)
	fmt.Print(root.Val, " ")
}

// ===== BFS =====
func levelOrder(root *TreeNode) {
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		fmt.Print(node.Val, " ")

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}

// ===== Height =====
func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := height(root.Left)
	right := height(root.Right)

	if left > right {
		return left + 1
	}
	return right + 1
}
