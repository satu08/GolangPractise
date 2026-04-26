package main

import (
	"fmt"
)

// 🔹 Print Matrix
func printMatrix(mat [][]int) {
	for _, row := range mat {
		fmt.Println(row)
	}
}

// 🔹 Spiral Traversal
func spiral(mat [][]int) []int {
	res := []int{}
	top, bottom := 0, len(mat)-1
	left, right := 0, len(mat[0])-1

	for top <= bottom && left <= right {

		for i := left; i <= right; i++ {
			res = append(res, mat[top][i])
		}
		top++

		for i := top; i <= bottom; i++ {
			res = append(res, mat[i][right])
		}
		right--

		if top <= bottom {
			for i := right; i >= left; i-- {
				res = append(res, mat[bottom][i])
			}
			bottom--
		}

		if left <= right {
			for i := bottom; i >= top; i-- {
				res = append(res, mat[i][left])
			}
			left++
		}
	}

	return res
}

// 🔹 Rotate Matrix 90° (clockwise)
func rotate(mat [][]int) {
	n := len(mat)

	// transpose
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			mat[i][j], mat[j][i] = mat[j][i], mat[i][j]
		}
	}

	// reverse rows
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			mat[i][j], mat[i][n-j-1] = mat[i][n-j-1], mat[i][j]
		}
	}
}

// 🔹 DFS (for grid problems)
func dfs(mat [][]int, i, j int) {
	rows, cols := len(mat), len(mat[0])

	if i < 0 || j < 0 || i >= rows || j >= cols || mat[i][j] == 0 {
		return
	}

	mat[i][j] = 0 // mark visited

	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for _, d := range dirs {
		dfs(mat, i+d[0], j+d[1])
	}
}

// 🔹 Count Islands (DFS)
func numIslands(mat [][]int) int {
	count := 0

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 1 {
				count++
				dfs(mat, i, j)
			}
		}
	}
	return count
}

// 🔹 MAIN
func main() {
	mat := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("Matrix:")
	printMatrix(mat)

	fmt.Println("\nSpiral Traversal:")
	fmt.Println(spiral(mat))

	fmt.Println("\nRotate 90°:")
	rotate(mat)
	printMatrix(mat)

	fmt.Println("\nNumber of Islands:")
	grid := [][]int{
		{1, 1, 0},
		{0, 1, 0},
		{1, 0, 1},
	}
	fmt.Println(numIslands(grid))
}
