package main

import "fmt"

func main() {

	graph := map[int][]int{
		0: {1, 2},
		1: {2},
		2: {0, 3},
		3: {3},
	}

	fmt.Print("DFS: ")
	visited := make(map[int]bool)
	dfs(0, graph, visited)

	fmt.Print("\nBFS: ")
	bfs(0, graph)
}

// ===== DFS =====
func dfs(node int, graph map[int][]int, visited map[int]bool) {
	visited[node] = true
	fmt.Print(node, " ")

	for _, nei := range graph[node] {
		if !visited[nei] {
			dfs(nei, graph, visited)
		}
	}
}

// ===== BFS =====
func bfs(start int, graph map[int][]int) {
	visited := make(map[int]bool)
	queue := []int{start}

	visited[start] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		fmt.Print(node, " ")

		for _, nei := range graph[node] {
			if !visited[nei] {
				visited[nei] = true
				queue = append(queue, nei)
			}
		}
	}
}
