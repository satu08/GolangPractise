package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	var root *TreeNode

	// Insert
	root = insert(root, 5)
	root = insert(root, 3)
	root = insert(root, 7)
	root = insert(root, 2)
	root = insert(root, 4)
	root = insert(root, 8)

	fmt.Print("Inorder (Sorted): ")
	inorder(root)

	// Search
	node := search(root, 4)
	if node != nil {
		fmt.Println("\nFound:", node.Val)
	}

	// Delete
	root = deleteNode(root, 3)
	fmt.Print("After Delete 3: ")
	inorder(root)

	// Validate
	fmt.Println("\nValid BST:", isValidBST(root, nil, nil))
}

// ===== Insert =====
func insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if val < root.Val {
		root.Left = insert(root.Left, val)
	} else {
		root.Right = insert(root.Right, val)
	}
	return root
}

// ===== Search =====
func search(root *TreeNode, key int) *TreeNode {
	if root == nil || root.Val == key {
		return root
	}

	if key < root.Val {
		return search(root.Left, key)
	}
	return search(root.Right, key)
}

// ===== Delete =====
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {

		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}

		minNode := findMin(root.Right)
		root.Val = minNode.Val
		root.Right = deleteNode(root.Right, minNode.Val)
	}
	return root
}

func findMin(root *TreeNode) *TreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}

// ===== Inorder =====
func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	fmt.Print(root.Val, " ")
	inorder(root.Right)
}

// ===== Validate =====
func isValidBST(root *TreeNode, min, max *int) bool {
	if root == nil {
		return true
	}

	if (min != nil && root.Val <= *min) ||
		(max != nil && root.Val >= *max) {
		return false
	}

	return isValidBST(root.Left, min, &root.Val) &&
		isValidBST(root.Right, &root.Val, max)
}
