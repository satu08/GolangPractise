Got it — you want a **clean, interview-ready array (slice) implementation in Go** with all common operations and one `main()` to test everything.

Let’s build a **mini dynamic array (like Java ArrayList / Python list)** 🚀

---

# 🧠 What we’ll implement

## Core Methods

* `Append`
* `Insert`
* `Delete`
* `Get`
* `Set`
* `Size`
* `Capacity`
* `IsEmpty`
* `Print`

---

# 🚀 Full Go Implementation

```go
package main

import "fmt"

type Array struct {
	data []int
	size int
}

// Constructor
func NewArray(cap int) *Array {
	return &Array{
		data: make([]int, cap),
		size: 0,
	}
}

// Resize (double capacity)
func (a *Array) resize() {
	newCap := len(a.data) * 2
	newData := make([]int, newCap)

	for i := 0; i < a.size; i++ {
		newData[i] = a.data[i]
	}
	a.data = newData
}

// Append element
func (a *Array) Append(val int) {
	if a.size == len(a.data) {
		a.resize()
	}
	a.data[a.size] = val
	a.size++
}

// Insert at index
func (a *Array) Insert(index int, val int) {
	if index < 0 || index > a.size {
		fmt.Println("Invalid index")
		return
	}

	if a.size == len(a.data) {
		a.resize()
	}

	for i := a.size; i > index; i-- {
		a.data[i] = a.data[i-1]
	}

	a.data[index] = val
	a.size++
}

// Delete at index
func (a *Array) Delete(index int) {
	if index < 0 || index >= a.size {
		fmt.Println("Invalid index")
		return
	}

	for i := index; i < a.size-1; i++ {
		a.data[i] = a.data[i+1]
	}

	a.size--
}

// Get element
func (a *Array) Get(index int) int {
	if index < 0 || index >= a.size {
		fmt.Println("Invalid index")
		return -1
	}
	return a.data[index]
}

// Set element
func (a *Array) Set(index int, val int) {
	if index < 0 || index >= a.size {
		fmt.Println("Invalid index")
		return
	}
	a.data[index] = val
}

// Size
func (a *Array) Size() int {
	return a.size
}

// Capacity
func (a *Array) Capacity() int {
	return len(a.data)
}

// IsEmpty
func (a *Array) IsEmpty() bool {
	return a.size == 0
}

// Print array
func (a *Array) Print() {
	for i := 0; i < a.size; i++ {
		fmt.Print(a.data[i], " ")
	}
	fmt.Println()
}

// MAIN FUNCTION (TEST ALL)
func main() {
	arr := NewArray(2)

	fmt.Println("Append:")
	arr.Append(10)
	arr.Append(20)
	arr.Append(30) // triggers resize
	arr.Print()

	fmt.Println("Insert at index 1:")
	arr.Insert(1, 15)
	arr.Print()

	fmt.Println("Delete index 2:")
	arr.Delete(2)
	arr.Print()

	fmt.Println("Get index 1:", arr.Get(1))

	fmt.Println("Set index 1 to 100:")
	arr.Set(1, 100)
	arr.Print()

	fmt.Println("Size:", arr.Size())
	fmt.Println("Capacity:", arr.Capacity())
	fmt.Println("IsEmpty:", arr.IsEmpty())
}
```

---

# 🧪 Output

```text
Append:
10 20 30
Insert at index 1:
10 15 20 30
Delete index 2:
10 15 30
Get index 1: 15
Set index 1 to 100:
10 100 30
Size: 3
Capacity: 4
IsEmpty: false
```

---

# 🧠 Interview Insights (VERY IMPORTANT)

## Why resize by 2?

👉 Amortized O(1) append

---

## Time Complexity

| Operation | Time           |
| --------- | -------------- |
| Append    | O(1) amortized |
| Insert    | O(n)           |
| Delete    | O(n)           |
| Get       | O(1)           |

---

## Space Complexity

```text
O(n)
```

---

# 🔥 Real-world mapping

This is how:

* Go slices work internally
* Java `ArrayList`
* Python `list`

---

# 🚀 If you want next level

I can extend this to:

* Generic array (Go generics)
* Thread-safe version
* Implement slice internals (capacity growth strategy)
* Compare with Go built-in slice behavior

Just say 👉 **"make it generic"** or **"deep dive slice internals"**



Perfect — let’s build this **clean + interview-ready** 🚀

You’ll get:

* ✅ Linear Search
* ✅ Binary Search (iterative + recursive)
* ✅ One `main()` to test everything

---

# 🧠 When to use what

| Algorithm     | When to use   |
| ------------- | ------------- |
| Linear Search | Unsorted data |
| Binary Search | Sorted data   |

---

# 🚀 Full Go Implementation

```go id="bcz3ux"
package main

import "fmt"

// 🔹 Linear Search (O(n))
func linearSearch(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

// 🔹 Binary Search - Iterative (O(log n))
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// 🔹 Binary Search - Recursive
func binarySearchRecursive(arr []int, left, right, target int) int {
	if left > right {
		return -1
	}

	mid := left + (right-left)/2

	if arr[mid] == target {
		return mid
	} else if arr[mid] < target {
		return binarySearchRecursive(arr, mid+1, right, target)
	} else {
		return binarySearchRecursive(arr, left, mid-1, target)
	}
}

func main() {

	// Unsorted array (for linear search)
	arr1 := []int{10, 5, 3, 8, 2}

	// Sorted array (for binary search)
	arr2 := []int{2, 3, 5, 8, 10}

	target := 8

	fmt.Println("🔹 Linear Search:")
	index1 := linearSearch(arr1, target)
	fmt.Println("Element found at index:", index1)

	fmt.Println("\n🔹 Binary Search (Iterative):")
	index2 := binarySearch(arr2, target)
	fmt.Println("Element found at index:", index2)

	fmt.Println("\n🔹 Binary Search (Recursive):")
	index3 := binarySearchRecursive(arr2, 0, len(arr2)-1, target)
	fmt.Println("Element found at index:", index3)
}
```

---

# 🧪 Output

```text
🔹 Linear Search:
Element found at index: 3

🔹 Binary Search (Iterative):
Element found at index: 3

🔹 Binary Search (Recursive):
Element found at index: 3
```

---

# 🧠 Key Interview Insights

## 🔹 Linear Search

* Works on **unsorted data**
* Simple but slow → O(n)

---

## 🔹 Binary Search (VERY IMPORTANT)

👉 Requirements:

* Array must be **sorted**

👉 Why fast?

* Reduces search space by half each time

---

## 🔥 Common Variations (Interview Favorites)

Be ready for:

* First occurrence
* Last occurrence
* Lower bound / Upper bound
* Search in rotated array 🔥
* Binary search on answer (very important)

---

# 🚀 Pro Tip (Senior level)

Always say this in interview:

> “If the problem has a monotonic condition, I’ll consider binary search on answer.”

---

# 🚀 Next?

If you want to go deeper:

👉 “binary search variations”
👉 “rotated array search”
👉 “binary search on answer problems”

These are **very high-frequency** 🔥


Perfect — this is a **core DSA pillar** and very important for interviews 🔥

I’ll give you:

* ✅ All major sorting algorithms
* ✅ When to use each
* ✅ One Go file with **all implementations + main()**
* ✅ Interview insights

---

# 🧠 Sorting Algorithms Overview

| Algorithm      | Time           | Stable | Use Case              |
| -------------- | -------------- | ------ | --------------------- |
| Bubble Sort    | O(n²)          | ✅      | Learning only         |
| Selection Sort | O(n²)          | ❌      | Rarely used           |
| Insertion Sort | O(n²)          | ✅      | Small / nearly sorted |
| Merge Sort     | O(n log n)     | ✅      | Stable sorting        |
| Quick Sort     | O(n log n) avg | ❌      | Fast in practice      |
| Heap Sort      | O(n log n)     | ❌      | Priority queues       |
| Counting Sort  | O(n+k)         | ✅      | Small range integers  |

---

# 🚀 Full Go Code (ALL SORTS + MAIN)

```go
package main

import (
	"fmt"
)

// 🔹 Bubble Sort
func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// 🔹 Selection Sort
func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

// 🔹 Insertion Sort
func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// 🔹 Merge Sort
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// 🔹 Quick Sort
func quickSort(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// 🔹 Heap Sort
func heapSort(arr []int) {
	n := len(arr)

	// build max heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// extract elements
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}

func heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

// 🔹 Counting Sort
func countingSort(arr []int) []int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}

	count := make([]int, max+1)

	for _, v := range arr {
		count[v]++
	}

	result := []int{}
	for i := 0; i <= max; i++ {
		for count[i] > 0 {
			result = append(result, i)
			count[i]--
		}
	}
	return result
}

// 🔹 MAIN
func main() {
	arr := []int{5, 2, 9, 1, 5, 6}

	fmt.Println("Original:", arr)

	// Bubble
	a1 := append([]int{}, arr...)
	bubbleSort(a1)
	fmt.Println("Bubble:", a1)

	// Selection
	a2 := append([]int{}, arr...)
	selectionSort(a2)
	fmt.Println("Selection:", a2)

	// Insertion
	a3 := append([]int{}, arr...)
	insertionSort(a3)
	fmt.Println("Insertion:", a3)

	// Merge
	a4 := mergeSort(arr)
	fmt.Println("Merge:", a4)

	// Quick
	a5 := append([]int{}, arr...)
	quickSort(a5, 0, len(a5)-1)
	fmt.Println("Quick:", a5)

	// Heap
	a6 := append([]int{}, arr...)
	heapSort(a6)
	fmt.Println("Heap:", a6)

	// Counting
	a7 := countingSort(arr)
	fmt.Println("Counting:", a7)
}
```

---

# 🧪 Output

```text
Original: [5 2 9 1 5 6]
Bubble: [1 2 5 5 6 9]
Selection: [1 2 5 5 6 9]
Insertion: [1 2 5 5 6 9]
Merge: [1 2 5 5 6 9]
Quick: [1 2 5 5 6 9]
Heap: [1 2 5 5 6 9]
Counting: [1 2 5 5 6 9]
```

---

# 🔥 Interview Insights (VERY IMPORTANT)

## 🧠 Most Asked

* Quick Sort
* Merge Sort
* Heap Sort

---

## 💡 When to use what

| Scenario            | Use            |
| ------------------- | -------------- |
| Nearly sorted       | Insertion Sort |
| Stable required     | Merge Sort     |
| Fast in practice    | Quick Sort     |
| Top K elements      | Heap           |
| Small range numbers | Counting Sort  |

---

# 🚀 Senior-Level Tip

Always say:

> “I’ll choose sorting based on constraints: stability, memory, and input distribution.”

---

# 🚀 Next?

If you want mastery:

👉 “quick sort dry run”
👉 “merge sort visual”
👉 “heap deep dive”
👉 “top 20 sorting interview questions”

This topic is **very high ROI** 🔥



Perfect — **Matrix is a high-frequency topic** (especially for BFS/DFS + grids) and also connects strongly to **graphs, DP, and recursion** 🔥

Let’s do it properly:

---

# 🧠 MATRIX — WHY IT EXISTS IN DSA

A matrix is just a **2D array**, but it models:

* Grids (maps, games, images)
* Graphs (implicit graphs)
* DP tables
* Real-world problems (like pathfinding)

---

# 🎯 WHERE MATRIX IS USED

* BFS / DFS on grids
* Shortest path problems
* Flood fill / islands
* Image processing
* DP (like unique paths)

---

# 🧠 CORE PATTERNS (VERY IMPORTANT)

## 1. 🔄 Traversal (Row / Column)

👉 Basic iteration

---

## 2. 🌀 Spiral Traversal

👉 Boundary-based traversal

---

## 3. 🔁 Rotation (90°)

👉 Transpose + reverse

---

## 4. 🌊 DFS / BFS on Grid (🔥 MOST IMPORTANT)

👉 Treat matrix as graph

---

## 5. 🧮 Prefix Sum (2D)

👉 Fast submatrix queries

---

# 🚀 FULL GO CODE (ALL IMPORTANT OPERATIONS)

```go
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

	dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}

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
```

---

# 🧪 Output

```text
Matrix:
[1 2 3]
[4 5 6]
[7 8 9]

Spiral Traversal:
[1 2 3 6 9 8 7 4 5]

Rotate 90°:
[7 4 1]
[8 5 2]
[9 6 3]

Number of Islands:
3
```

---

# 🔥 Interview Patterns (VERY IMPORTANT)

## 🟢 Pattern 1: Direction Array

```go
dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
```

👉 Used in:

* BFS
* DFS
* Grid traversal

---

## 🟢 Pattern 2: Boundary Check

```go
if i<0 || j<0 || i>=rows || j>=cols
```

---

## 🟢 Pattern 3: Mark Visited

```go
mat[i][j] = 0
```

OR use `visited[][]`

---

# 🎯 MOST ASKED MATRIX PROBLEMS

You should practice:

* Number of Islands 🔥
* Rotting Oranges 🔥
* Spiral Matrix
* Rotate Image
* Word Search 🔥
* Shortest Path in Binary Matrix

---

# 🚀 Senior-Level Insight

👉 Matrix = **Graph in disguise**

Always think:

* Nodes = cells
* Edges = directions

---

# 🚀 Next Step

Say:

👉 **"give matrix problems with patterns"**
👉 **"BFS vs DFS deep dive"**
👉 **"graph next"**

You’re progressing exactly like a **top candidate** 🔥




Perfect — **Hashing is one of the highest ROI topics in interviews** 🔥
(You’ll solve ~30–40% problems using this)

Let’s go **deep + practical** like an interviewer expects.

---

# 🧠 HASHING — WHY IT EXISTS

👉 To get **O(1) lookup time**

Instead of:

```text
search in array → O(n)
```

We use:

```text
hash map → O(1)
```

---

# 🎯 WHERE HASHING IS USED

* Frequency counting
* Duplicate detection
* Subarray problems
* Caching (like LRU)
* Strings (anagrams, patterns)

---

# 🧠 CORE IDEA

👉 Map key → value

```go
mp := map[int]int{}
```

---

# 🔥 MOST IMPORTANT PATTERNS

---

## 🟢 1. Frequency Count

### 👉 Use when:

* “count occurrences”
* “most frequent”

```go
func frequency(arr []int) map[int]int {
	mp := make(map[int]int)

	for _, v := range arr {
		mp[v]++
	}
	return mp
}
```

---

## 🟢 2. Two Sum Pattern

### 👉 Use when:

* “find pair with sum = target”

```go
func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)

	for i, v := range nums {
		if idx, ok := mp[target-v]; ok {
			return []int{idx, i}
		}
		mp[v] = i
	}
	return nil
}
```

---

## 🟢 3. Set for Duplicates

```go
func containsDuplicate(nums []int) bool {
	set := make(map[int]bool)

	for _, v := range nums {
		if set[v] {
			return true
		}
		set[v] = true
	}
	return false
}
```

---

## 🟢 4. Prefix Sum + HashMap (🔥 VERY IMPORTANT)

### 👉 Use when:

* “subarray sum = k”

```go
func subarraySum(nums []int, k int) int {
	mp := map[int]int{0: 1}
	sum, count := 0, 0

	for _, v := range nums {
		sum += v

		if c, ok := mp[sum-k]; ok {
			count += c
		}

		mp[sum]++
	}
	return count
}
```

---

## 🟢 5. Longest Subarray Pattern

```go
func longestSubarray(nums []int, k int) int {
	mp := make(map[int]int)
	sum, maxLen := 0, 0

	for i, v := range nums {
		sum += v

		if sum == k {
			maxLen = i + 1
		}

		if idx, ok := mp[sum-k]; ok {
			if i-idx > maxLen {
				maxLen = i - idx
			}
		}

		if _, ok := mp[sum]; !ok {
			mp[sum] = i
		}
	}
	return maxLen
}
```

---

# 🚀 FULL GO PROGRAM (ALL IN ONE)

```go id="h8u3kz"
package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 2, 1, 4}

	// 🔹 Frequency
	freq := make(map[int]int)
	for _, v := range arr {
		freq[v]++
	}
	fmt.Println("Frequency:", freq)

	// 🔹 Two Sum
	fmt.Println("Two Sum:", twoSum(arr, 5))

	// 🔹 Duplicate
	fmt.Println("Contains Duplicate:", containsDuplicate(arr))

	// 🔹 Subarray Sum
	fmt.Println("Subarray Sum =", subarraySum(arr, 3))
}

// helper functions
func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for i, v := range nums {
		if idx, ok := mp[target-v]; ok {
			return []int{idx, i}
		}
		mp[v] = i
	}
	return nil
}

func containsDuplicate(nums []int) bool {
	set := make(map[int]bool)
	for _, v := range nums {
		if set[v] {
			return true
		}
		set[v] = true
	}
	return false
}

func subarraySum(nums []int, k int) int {
	mp := map[int]int{0: 1}
	sum, count := 0, 0

	for _, v := range nums {
		sum += v
		if c, ok := mp[sum-k]; ok {
			count += c
		}
		mp[sum]++
	}
	return count
}
```

---

# 🧪 Output

```text id="6hq4bd"
Frequency: map[1:2 2:2 3:1 4:1]
Two Sum: [1 2]
Contains Duplicate: true
Subarray Sum = 3
```

---

# 🔥 Interview Signals (VERY IMPORTANT)

If you see:

* “count”
* “frequency”
* “duplicate”
* “pair”
* “subarray sum”

👉 **Think HASHING immediately**

---

# 🚀 Senior-Level Insights

## ⚠️ Collision (conceptual)

Different keys → same hash
Handled internally in Go maps

---

## ⚡ Tradeoff

| Advantage   | Disadvantage |
| ----------- | ------------ |
| O(1) lookup | Extra space  |

---

# 🚀 Next Step

Now best progression is:

👉 **Linked List (for LRU, stack, queue)**
or
👉 **Stack (monotonic stack problems 🔥)**

---

Just say:

👉 “next linked list”
👉 “stack deep dive”
👉 “give hashing problems”

We’ll go to **next level 🚀**




Great — **Strings is one of the most asked topics in interviews** 🔥
(and very important for backend + parsing use cases like your work)

Let’s go **deep + pattern-based (not just syntax)**

---

# 🧠 STRINGS — WHY IT EXISTS

👉 String = **array of characters**

Used in:

* Parsing APIs / logs
* Protocols (BACnet, HTTP)
* Text processing
* Validation problems

---

# ⚠️ IMPORTANT (Go-specific)

Strings in Go are:

* **immutable**
* UTF-8 encoded

👉 So modifying string = create new string

---

# 🎯 CORE PATTERNS (VERY IMPORTANT)

---

## 🟢 1. Two Pointers

👉 Used for:

* reverse
* palindrome

---

### ✅ Example: Palindrome

```go id="gqrf0i"
func isPalindrome(s string) bool {
	i, j := 0, len(s)-1

	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
```

---

## 🟢 2. Sliding Window 🔥

👉 Used for:

* substring problems
* longest/shortest

---

### ✅ Example: Longest Unique Substring

```go id="dx5i2o"
func longestUnique(s string) int {
	set := make(map[byte]bool)
	left, maxLen := 0, 0

	for right := 0; right < len(s); right++ {

		for set[s[right]] {
			delete(set, s[left])
			left++
		}

		set[s[right]] = true

		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}
```

---

## 🟢 3. Frequency Count (Hashing)

👉 Used for:

* anagrams
* character count

---

### ✅ Example: Valid Anagram

```go id="kyzw1l"
func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		count[s[i]]++
		count[t[i]]--
	}

	for _, v := range count {
		if v != 0 {
			return false
		}
	}
	return true
}
```

---

## 🟢 4. String Building

👉 Since string immutable → use builder

---

### ✅ Example

```go id="plkqg2"
import "strings"

var sb strings.Builder
sb.WriteString("hello")
sb.WriteString("world")

result := sb.String()
```

---

## 🟢 5. Pattern Matching (Advanced)

* KMP
* Rabin-Karp

👉 Rare but asked in top companies

---

# 🚀 FULL GO PROGRAM (ALL IN ONE)

```go id="a7mq2h"
package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "abcabcbb"

	fmt.Println("Longest Unique:", longestUnique(s))

	fmt.Println("Palindrome (racecar):", isPalindrome("racecar"))

	fmt.Println("Anagram (listen, silent):", isAnagram("listen", "silent"))

	// String Builder
	var sb strings.Builder
	sb.WriteString("Go ")
	sb.WriteString("Lang")
	fmt.Println("Builder:", sb.String())
}

// 🔹 Longest substring
func longestUnique(s string) int {
	set := make(map[byte]bool)
	left, maxLen := 0, 0

	for right := 0; right < len(s); right++ {
		for set[s[right]] {
			delete(set, s[left])
			left++
		}
		set[s[right]] = true
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}

// 🔹 Palindrome
func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// 🔹 Anagram
func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		count[s[i]]++
		count[t[i]]--
	}

	for _, v := range count {
		if v != 0 {
			return false
		}
	}
	return true
}
```

---

# 🧪 Output

```text
Longest Unique: 3
Palindrome (racecar): true
Anagram (listen, silent): true
Builder: Go Lang
```

---

# 🔥 MOST ASKED STRING PROBLEMS

You MUST practice:

* Longest Substring Without Repeating 🔥
* Minimum Window Substring 🔥
* Valid Anagram
* Group Anagrams
* String to Integer (atoi)
* KMP pattern matching

---

# 🚀 Senior-Level Insights

## ⚠️ ASCII vs Unicode

```go id="8mk5cg"
for _, ch := range s { // correct for unicode
```

👉 Don’t use `s[i]` for unicode strings

---

## ⚡ Performance Tip

Use:

* `strings.Builder`
* `[]byte`

👉 instead of repeated string concatenation

---

# 🚀 Interview Signals

If you see:

* substring
* palindrome
* anagram
* pattern

👉 Think:

```text
Sliding Window + Hashing
```

---

# 🚀 Next Step

Best progression now:

👉 **Stack (very important for strings + NGE)**
👉 **Linked List (LRU + system design)**
👉 **Tree (must for senior roles)**

---

Just say:

👉 “stack deep dive”
👉 “linked list next”
👉 “string problems set”

You’re building **very strong DSA foundation 🔥**




Perfect — **Linked List is a must-know topic**, especially for:

* LRU Cache (you already saw 🔥)
* System design internals
* Pointer manipulation questions

Let’s go **deep + practical + interview-ready**

---

# 🧠 LINKED LIST — WHY IT EXISTS

👉 Unlike arrays:

* No contiguous memory needed
* Dynamic size
* Easy insert/delete

---

# 📦 Types of Linked List

## 🔹 Singly Linked List

```text
A → B → C → D
```

## 🔹 Doubly Linked List

```text
A ⇄ B ⇄ C ⇄ D
```

## 🔹 Circular Linked List

```text
A → B → C → A
```

---

# 🎯 WHERE USED

* LRU Cache (DLL + HashMap)
* Stack / Queue
* Memory management
* Graph adjacency list

---

# 🧠 CORE PATTERNS (VERY IMPORTANT)

---

## 🟢 1. Traversal

👉 Move pointer until nil

---

## 🟢 2. Reverse Linked List 🔥

👉 Most asked

---

## 🟢 3. Fast & Slow Pointer 🔥

👉 Detect:

* cycle
* middle node

---

## 🟢 4. Dummy Node Pattern

👉 Avoid edge cases (VERY IMPORTANT)

---

# 🚀 FULL GO IMPLEMENTATION (ALL IN ONE)

```go id="s6kgzq"
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
```

---

# 🧪 Output

```text
Original List:
1 -> 2 -> 3 -> 4 -> nil

Reversed:
4 -> 3 -> 2 -> 1 -> nil

Middle: 2

After Delete 3:
4 -> 2 -> 1 -> nil

Has Cycle: true
```

---

# 🔥 Interview Patterns (VERY IMPORTANT)

## 🟢 Reverse Pattern

```go
next := curr.Next
curr.Next = prev
```

---

## 🟢 Fast & Slow

```go
slow = slow.Next
fast = fast.Next.Next
```

---

## 🟢 Dummy Node

```go
dummy := &ListNode{Next: head}
```

👉 Avoids edge cases (head deletion)

---

# 🎯 MOST ASKED PROBLEMS

You MUST practice:

* Reverse Linked List 🔥
* Linked List Cycle 🔥
* Merge Two Sorted Lists
* Remove Nth Node from End 🔥
* Intersection of Two Lists
* LRU Cache 🔥

---

# 🚀 Senior-Level Insights

## ⚠️ Why Linked List in real systems?

* Memory efficient insert/delete
* Used in:

  * OS scheduling
  * caches
  * allocators

---

## ⚡ Tradeoff

| Advantage          | Disadvantage |
| ------------------ | ------------ |
| O(1) insert/delete | O(n) access  |

---

# 🚀 Next Step

Best progression now:

👉 **Stack (very important + monotonic stack)**
👉 **Queue / Deque**
👉 **Tree (VERY IMPORTANT for interviews)**

---

Just say:

👉 “stack deep dive”
👉 “queue next”
👉 “tree full guide”

You’re building **strong interview foundation 🔥**




Perfect — **Stack is one of the most powerful patterns in DSA** 🔥
(especially for **monotonic stack, parsing, and recursion replacement**)

Let’s go **deep + pattern-based + Go implementation**

---

# 🧠 STACK — WHY IT EXISTS

👉 Follows **LIFO (Last In First Out)**

```text
Push → Push → Pop
```

---

# 🎯 WHERE STACK IS USED

* Expression evaluation
* Undo/Redo systems
* Function calls (call stack)
* Parsing (compiler)
* Monotonic stack problems 🔥

---

# 🧠 CORE OPERATIONS

| Operation | Meaning     |
| --------- | ----------- |
| Push      | Add element |
| Pop       | Remove top  |
| Peek      | Top element |
| IsEmpty   | Check empty |

---

# 🚀 STACK IMPLEMENTATION IN GO

```go id="u2c6tn"
package main

import "fmt"

type Stack struct {
	data []int
}

// Push
func (s *Stack) Push(val int) {
	s.data = append(s.data, val)
}

// Pop
func (s *Stack) Pop() int {
	if len(s.data) == 0 {
		return -1
	}
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val
}

// Peek
func (s *Stack) Peek() int {
	if len(s.data) == 0 {
		return -1
	}
	return s.data[len(s.data)-1]
}

// IsEmpty
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}
```

---

# 🔥 IMPORTANT PATTERNS

---

## 🟢 1. Valid Parentheses (Classic)

👉 Matching brackets

```go id="thzprg"
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
```

---

## 🟢 2. Next Greater Element 🔥

👉 Monotonic stack

```go id="cwyzy9"
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
```

---

## 🟢 3. Largest Rectangle in Histogram 🔥

👉 Hard but very common

---

## 🟢 4. Stock Span Problem

---

## 🟢 5. Reverse using Stack

---

# 🚀 FULL PROGRAM (ALL IN ONE)

```go id="7kh21r"
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
```

---

# 🧪 Output

```text id="h4yq6z"
Pop: 30
Peek: 20
Valid: true
Next Greater: [3 4 4 -1]
```

---

# 🔥 MOST IMPORTANT INTERVIEW PATTERN

## 👉 Monotonic Stack

### Types:

* Increasing stack
* Decreasing stack

---

## 🧠 Use when:

* Next greater/smaller
* Previous greater/smaller
* Range queries

---

# 🎯 HIGH-FREQUENCY PROBLEMS

* Next Greater Element 🔥
* Stock Span 🔥
* Largest Rectangle 🔥
* Trapping Rain Water 🔥
* Valid Parentheses

---

# 🚀 Senior-Level Insight

👉 Stack is used internally in:

* recursion
* compilers
* expression evaluation

---

# 🚀 Next Step

Now best progression:

👉 **Queue / Deque (important for BFS)**
👉 **Tree (VERY VERY IMPORTANT)**
👉 **Graph (advanced level)**

---

Just say:

👉 “queue next”
👉 “deque deep dive”
👉 “tree start”

You’re progressing like a **top interview candidate 🔥**




Perfect — **Queue is fundamental for BFS, scheduling, streaming systems** 🔥
(and very important for system design + graph problems)

Let’s go **clear + practical + interview-focused**

---

# 🧠 QUEUE — WHY IT EXISTS

👉 Follows **FIFO (First In First Out)**

```text
Enqueue → Enqueue → Dequeue
```

---

# 🎯 WHERE QUEUE IS USED

* BFS (graphs, trees) 🔥
* Task scheduling
* Rate limiting
* Streaming systems
* OS process queues

---

# 🧠 CORE OPERATIONS

| Operation | Meaning                |
| --------- | ---------------------- |
| Enqueue   | Add element (rear)     |
| Dequeue   | Remove element (front) |
| Front     | Peek first             |
| IsEmpty   | Check empty            |

---

# 🚀 QUEUE IMPLEMENTATION IN GO (Slice-based)

```go id="6xk8xa"
package main

import "fmt"

type Queue struct {
	data []int
}

// Enqueue
func (q *Queue) Enqueue(val int) {
	q.data = append(q.data, val)
}

// Dequeue
func (q *Queue) Dequeue() int {
	if len(q.data) == 0 {
		return -1
	}
	val := q.data[0]
	q.data = q.data[1:]
	return val
}

// Front
func (q *Queue) Front() int {
	if len(q.data) == 0 {
		return -1
	}
	return q.data[0]
}

// IsEmpty
func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}
```

---

# ⚠️ Important Insight

👉 `q.data = q.data[1:]`
This **does not free memory immediately** → can cause memory issues

👉 For production:

* use circular queue OR
* use `container/list`

---

# 🔥 IMPORTANT PATTERNS

---

## 🟢 1. BFS (VERY IMPORTANT 🔥)

👉 Level order traversal

---

### ✅ Example: BFS (Tree)

```go id="e6kn0z"
func bfs(root *TreeNode) {
	if root == nil {
		return
	}

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
```

---

## 🟢 2. Level Order Traversal

👉 Very common interview question

---

## 🟢 3. Sliding Window Maximum (Deque + Queue hybrid)

---

## 🟢 4. Circular Queue

👉 Avoid memory waste

---

# 🚀 FULL PROGRAM (ALL IN ONE)

```go id="q0e9g2"
package main

import "fmt"

// Queue
type Queue struct {
	data []int
}

func (q *Queue) Enqueue(val int) {
	q.data = append(q.data, val)
}

func (q *Queue) Dequeue() int {
	if len(q.data) == 0 {
		return -1
	}
	val := q.data[0]
	q.data = q.data[1:]
	return val
}

func (q *Queue) Front() int {
	if len(q.data) == 0 {
		return -1
	}
	return q.data[0]
}

// Tree Node (for BFS demo)
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BFS
func bfs(root *TreeNode) {
	if root == nil {
		return
	}

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

func main() {

	// 🔹 Queue operations
	q := Queue{}
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	fmt.Println("Dequeue:", q.Dequeue())
	fmt.Println("Front:", q.Front())

	// 🔹 BFS example
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}

	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	fmt.Print("BFS: ")
	bfs(root)
}
```

---

# 🧪 Output

```text
Dequeue: 10
Front: 20
BFS: 1 2 3 4 5
```

---

# 🔥 INTERVIEW PATTERNS (VERY IMPORTANT)

## 🟢 Queue + BFS

👉 Use when:

* shortest path
* level traversal
* grid problems

---

## 🟢 Multi-source BFS 🔥

Used in:

* Rotting Oranges
* Distance problems

---

## 🟢 Circular Queue

👉 Avoid memory leaks

---

# 🎯 MOST ASKED PROBLEMS

You MUST practice:

* Level Order Traversal 🔥
* Rotting Oranges 🔥
* Sliding Window Maximum 🔥
* Implement Queue using Stack
* Circular Queue

---

# 🚀 Senior-Level Insight

👉 Queue is backbone of:

* BFS
* Event-driven systems
* Message queues (Kafka, RabbitMQ)

---

# 🚀 Next Step

Best progression now:

👉 **Deque (important for sliding window)**
👉 **Tree (VERY VERY IMPORTANT)**
👉 **Graph (advanced + BFS/DFS)**

---

Just say:

👉 “deque deep dive”
👉 “tree start”
👉 “graph next”

You’re building **strong DSA fundamentals 🔥**


Perfect — **Queue is fundamental for BFS, scheduling, streaming systems** 🔥
(and very important for system design + graph problems)

Let’s go **clear + practical + interview-focused**

---

# 🧠 QUEUE — WHY IT EXISTS

👉 Follows **FIFO (First In First Out)**

```text
Enqueue → Enqueue → Dequeue
```

---

# 🎯 WHERE QUEUE IS USED

* BFS (graphs, trees) 🔥
* Task scheduling
* Rate limiting
* Streaming systems
* OS process queues

---

# 🧠 CORE OPERATIONS

| Operation | Meaning                |
| --------- | ---------------------- |
| Enqueue   | Add element (rear)     |
| Dequeue   | Remove element (front) |
| Front     | Peek first             |
| IsEmpty   | Check empty            |

---

# 🚀 QUEUE IMPLEMENTATION IN GO (Slice-based)

```go id="6xk8xa"
package main

import "fmt"

type Queue struct {
	data []int
}

// Enqueue
func (q *Queue) Enqueue(val int) {
	q.data = append(q.data, val)
}

// Dequeue
func (q *Queue) Dequeue() int {
	if len(q.data) == 0 {
		return -1
	}
	val := q.data[0]
	q.data = q.data[1:]
	return val
}

// Front
func (q *Queue) Front() int {
	if len(q.data) == 0 {
		return -1
	}
	return q.data[0]
}

// IsEmpty
func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}
```

---

# ⚠️ Important Insight

👉 `q.data = q.data[1:]`
This **does not free memory immediately** → can cause memory issues

👉 For production:

* use circular queue OR
* use `container/list`

---

# 🔥 IMPORTANT PATTERNS

---

## 🟢 1. BFS (VERY IMPORTANT 🔥)

👉 Level order traversal

---

### ✅ Example: BFS (Tree)

```go id="e6kn0z"
func bfs(root *TreeNode) {
	if root == nil {
		return
	}

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
```

---

## 🟢 2. Level Order Traversal

👉 Very common interview question

---

## 🟢 3. Sliding Window Maximum (Deque + Queue hybrid)

---

## 🟢 4. Circular Queue

👉 Avoid memory waste

---

# 🚀 FULL PROGRAM (ALL IN ONE)

```go id="q0e9g2"
package main

import "fmt"

// Queue
type Queue struct {
	data []int
}

func (q *Queue) Enqueue(val int) {
	q.data = append(q.data, val)
}

func (q *Queue) Dequeue() int {
	if len(q.data) == 0 {
		return -1
	}
	val := q.data[0]
	q.data = q.data[1:]
	return val
}

func (q *Queue) Front() int {
	if len(q.data) == 0 {
		return -1
	}
	return q.data[0]
}

// Tree Node (for BFS demo)
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BFS
func bfs(root *TreeNode) {
	if root == nil {
		return
	}

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

func main() {

	// 🔹 Queue operations
	q := Queue{}
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	fmt.Println("Dequeue:", q.Dequeue())
	fmt.Println("Front:", q.Front())

	// 🔹 BFS example
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}

	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	fmt.Print("BFS: ")
	bfs(root)
}
```

---

# 🧪 Output

```text
Dequeue: 10
Front: 20
BFS: 1 2 3 4 5
```

---

# 🔥 INTERVIEW PATTERNS (VERY IMPORTANT)

## 🟢 Queue + BFS

👉 Use when:

* shortest path
* level traversal
* grid problems

---

## 🟢 Multi-source BFS 🔥

Used in:

* Rotting Oranges
* Distance problems

---

## 🟢 Circular Queue

👉 Avoid memory leaks

---

# 🎯 MOST ASKED PROBLEMS

You MUST practice:

* Level Order Traversal 🔥
* Rotting Oranges 🔥
* Sliding Window Maximum 🔥
* Implement Queue using Stack
* Circular Queue

---

# 🚀 Senior-Level Insight

👉 Queue is backbone of:

* BFS
* Event-driven systems
* Message queues (Kafka, RabbitMQ)

---

# 🚀 Next Step

Best progression now:

👉 **Deque (important for sliding window)**
👉 **Tree (VERY VERY IMPORTANT)**
👉 **Graph (advanced + BFS/DFS)**

---

Just say:

👉 “deque deep dive”
👉 “tree start”
👉 “graph next”

You’re building **strong DSA fundamentals 🔥**




Perfect — **Deque (Double Ended Queue)** is a **very high-value topic** 🔥
(It appears in **sliding window, monotonic queue, BFS optimizations**)

Let’s go **deep + pattern-based + Go implementation**

---

# 🧠 DEQUE — WHY IT EXISTS

👉 A **Deque allows insertion/removal from BOTH ends**

```text
Front <--> [Deque] <--> Rear
```

---

# 🎯 WHERE DEQUE IS USED

* Sliding Window Maximum 🔥
* Monotonic Queue 🔥
* BFS optimization (0-1 BFS)
* LRU Cache (internally)
* Real-time streaming problems

---

# 🧠 OPERATIONS

| Operation | Meaning         |
| --------- | --------------- |
| PushFront | Insert at front |
| PushBack  | Insert at rear  |
| PopFront  | Remove front    |
| PopBack   | Remove rear     |
| Front     | Peek front      |
| Back      | Peek rear       |

---

# 🚀 DEQUE IMPLEMENTATION IN GO (using slice)

```go id="gcf6gj"
package main

import "fmt"

type Deque struct {
	data []int
}

// Push Front
func (d *Deque) PushFront(val int) {
	d.data = append([]int{val}, d.data...)
}

// Push Back
func (d *Deque) PushBack(val int) {
	d.data = append(d.data, val)
}

// Pop Front
func (d *Deque) PopFront() int {
	if len(d.data) == 0 {
		return -1
	}
	val := d.data[0]
	d.data = d.data[1:]
	return val
}

// Pop Back
func (d *Deque) PopBack() int {
	if len(d.data) == 0 {
		return -1
	}
	val := d.data[len(d.data)-1]
	d.data = d.data[:len(d.data)-1]
	return val
}

// Front
func (d *Deque) Front() int {
	if len(d.data) == 0 {
		return -1
	}
	return d.data[0]
}

// Back
func (d *Deque) Back() int {
	if len(d.data) == 0 {
		return -1
	}
	return d.data[len(d.data)-1]
}
```

---

# ⚠️ Important Note

👉 `PushFront` is **O(n)** (shift operation)

👉 In real systems:

* Use **circular buffer**
* Or Go’s `container/list`

---

# 🔥 MOST IMPORTANT PATTERN

# 🟢 Monotonic Queue (🔥 VERY IMPORTANT)

👉 Used in:

* Sliding Window Maximum
* Optimization problems

---

## 🚀 Example: Sliding Window Maximum

### 🧠 Idea

* Maintain decreasing deque
* Front always holds max

---

### ✅ Go Code

```go id="fsu9xv"
func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	deque := []int{} // store indices

	for i := 0; i < len(nums); i++ {

		// remove out-of-window
		if len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}

		// maintain decreasing order
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
```

---

# 🧪 Example

```text
Input: [1,3,-1,-3,5,3,6,7], k = 3
Output: [3,3,5,5,6,7]
```

---

# 🚀 FULL PROGRAM (ALL IN ONE)

```go id="68bq6p"
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
	arr := []int{1,3,-1,-3,5,3,6,7}
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
```

---

# 🧪 Output

```text
Front: 5
Back: 20
After PopFront: [10 20]
Sliding Window: [3 3 5 5 6 7]
```

---

# 🔥 INTERVIEW PATTERNS

## 🟢 Deque + Sliding Window

👉 Max/Min in window

---

## 🟢 Monotonic Queue

👉 Maintain order (increasing/decreasing)

---

## 🟢 0-1 BFS

👉 Graph optimization using deque

---

# 🎯 MOST ASKED PROBLEMS

* Sliding Window Maximum 🔥
* Shortest Subarray with Sum ≥ K 🔥
* 0-1 BFS 🔥
* Maximum in subarrays

---

# 🚀 Senior-Level Insight

👉 Deque = **Queue + Stack combined**

Used in:

* real-time analytics
* streaming systems
* optimization problems

---

# 🚀 Next Step

Now we move into **CORE INTERVIEW ZONE**:

👉 **Tree (VERY VERY IMPORTANT 🔥)**
👉 **Binary Search Tree**
👉 **Graph (final boss)**

---

Just say:

👉 **“tree start”**

You’re about to enter **top interview topics 🚀**





Perfect — now you’re entering the **CORE of DSA interviews** 🔥
👉 **Trees are asked in almost every serious interview**

Let’s go **deep + pattern-based + Go implementation**

---

# 🌲 TREE — WHY IT EXISTS

👉 To represent **hierarchical data**

Examples:

* File systems
* Org charts
* DOM (HTML)
* Databases (indexes)

---

# 🧠 BASIC STRUCTURE

## Binary Tree

```text
        1
       / \
      2   3
     / \
    4   5
```

---

# 🎯 WHERE TREES ARE USED

* Searching (BST)
* Heaps
* Graph traversal
* AI / decision trees
* System design (indexes, caching)

---

# 🧠 CORE PATTERNS (VERY IMPORTANT)

---

## 🟢 1. DFS (Depth First Search) 🔥

Types:

* Preorder
* Inorder
* Postorder

---

## 🟢 2. BFS (Level Order) 🔥

👉 Uses Queue

---

## 🟢 3. Recursion (MOST IMPORTANT)

👉 Trees = recursion naturally

---

## 🟢 4. Tree DP

👉 Solve bottom-up

---

# 🚀 TREE IMPLEMENTATION IN GO

```go id="j5h9ra"
package main

import "fmt"

// Tree Node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
```

---

# 🔥 DFS TRAVERSALS

---

## ✅ Preorder (Root → Left → Right)

```go id="p9x3px"
func preorder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Val, " ")
	preorder(root.Left)
	preorder(root.Right)
}
```

---

## ✅ Inorder (Left → Root → Right)

```go id="o2i9k3"
func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	fmt.Print(root.Val, " ")
	inorder(root.Right)
}
```

---

## ✅ Postorder (Left → Right → Root)

```go id="z6b8wr"
func postorder(root *TreeNode) {
	if root == nil {
		return
	}
	postorder(root.Left)
	postorder(root.Right)
	fmt.Print(root.Val, " ")
}
```

---

# 🔥 BFS (LEVEL ORDER)

```go id="5q7r7x"
func levelOrder(root *TreeNode) {
	if root == nil {
		return
	}

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
```

---

# 🔥 IMPORTANT PROBLEMS

---

## 🟢 1. Height of Tree

```go id="m9h4rx"
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
```

---

## 🟢 2. Check Balanced Tree

---

## 🟢 3. Diameter of Tree 🔥

---

## 🟢 4. Lowest Common Ancestor 🔥

---

# 🚀 FULL PROGRAM (ALL IN ONE)

```go id="h7c9mf"
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
```

---

# 🧪 Output

```text
Preorder: 1 2 4 5 3
Inorder: 4 2 5 1 3
Postorder: 4 5 2 3 1
Level Order: 1 2 3 4 5
Height: 3
```

---

# 🔥 INTERVIEW PATTERNS (VERY IMPORTANT)

## 🟢 DFS Pattern

```go id="v6t1q2"
if root == nil { return }
```

---

## 🟢 BFS Pattern

```go id="m4z9xy"
queue := []*TreeNode{root}
```

---

## 🟢 Divide & Conquer

```go id="v0c9js"
left := solve(root.Left)
right := solve(root.Right)
```

---

# 🎯 MOST ASKED TREE PROBLEMS

You MUST practice:

* Level Order Traversal 🔥
* Diameter of Tree 🔥
* Lowest Common Ancestor 🔥
* Balanced Binary Tree
* Path Sum

---

# 🚀 Senior-Level Insight

👉 Tree = **Graph with constraints**

👉 Many problems = DFS/BFS + recursion

---

# 🚀 Next Step (VERY IMPORTANT)

Now go to:

👉 **Binary Search Tree (BST)**
👉 **Heap (priority queue)**
👉 **Graph (FINAL boss)**

---

Just say:

👉 **“bst next”**
👉 **“heap deep dive”**
👉 **“graph start”**

You’re now entering **top interview territory 🚀🔥**








Perfect — **Binary Search Tree (BST)** is one of the **most important interview topics** 🔥
(used in databases, indexing, searching systems)

Let’s go **deep + pattern-based + Go implementation**

---

# 🌲 BST — WHY IT EXISTS

👉 To perform **fast search, insert, delete**

---

## 🧠 Property (VERY IMPORTANT)

```text
Left < Root < Right
```

Example:

```text
        5
       / \
      3   7
     / \   \
    2   4   8
```

---

# 🎯 WHERE BST IS USED

* Databases (indexing)
* File systems
* Search engines
* Ordered data storage

---

# 🧠 CORE OPERATIONS (INTERVIEW MUST)

---

## 🟢 1. Search

👉 Like binary search

```go
func search(root *TreeNode, key int) *TreeNode {
	if root == nil || root.Val == key {
		return root
	}

	if key < root.Val {
		return search(root.Left, key)
	}
	return search(root.Right, key)
}
```

---

## 🟢 2. Insert

```go
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
```

---

## 🟢 3. Delete 🔥 (IMPORTANT)

3 cases:

1. No child
2. One child
3. Two children (use inorder successor)

```go
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		// found node

		// no left child
		if root.Left == nil {
			return root.Right
		}

		// no right child
		if root.Right == nil {
			return root.Left
		}

		// two children
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
```

---

# 🔥 IMPORTANT PATTERNS

---

## 🟢 4. Validate BST 🔥

```go
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
```

---

## 🟢 5. Kth Smallest 🔥

👉 Inorder traversal

---

# 🚀 FULL GO PROGRAM

```go
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
```

---

# 🧪 Output

```text
Inorder (Sorted): 2 3 4 5 7 8
Found: 4
After Delete 3: 2 4 5 7 8
Valid BST: true
```

---

# 🔥 INTERVIEW PATTERNS

## 🟢 Inorder = Sorted Order 🔥

---

## 🟢 BST Search = O(log n)

(average case)

---

## 🟢 Delete = tricky (3 cases)

---

# 🎯 MOST ASKED BST PROBLEMS

* Validate BST 🔥
* Kth smallest 🔥
* LCA in BST 🔥
* Convert sorted array → BST
* Range sum BST

---

# 🚀 Senior-Level Insight

👉 BST becomes slow if unbalanced
👉 So we use:

* AVL Tree
* Red-Black Tree (used in systems)

---

# 🚀 Next Step

Now we go to:

👉 **Heap (Priority Queue) 🔥**
👉 **Graph (FINAL boss)**
👉 **Greedy / DP (advanced)**

---

Just say:

👉 **“heap deep dive”**
👉 **“graph start”**

You’re now at **core interview level 🚀🔥**



Perfect — **Heap (Priority Queue)** is 🔥 **super important for interviews**
Used in:

* Top K problems
* Scheduling
* Dijkstra (graphs)
* Streaming data

Let’s go **deep + practical + Go-focused**

---

# 🧠 HEAP — WHY IT EXISTS

👉 Efficiently get:

* **Minimum** (Min Heap)
* **Maximum** (Max Heap)

---

## 🎯 Property

### Min Heap

```text
Parent ≤ Children
```

### Max Heap

```text
Parent ≥ Children
```

---

## 📦 Representation (Array)

```text
Index:  0  1  2  3  4  5
Value: [1, 3, 5, 7, 9, 6]
```

Relationships:

```text
Left  = 2*i + 1
Right = 2*i + 2
Parent = (i-1)/2
```

---

# 🚀 HEAP IN GO (IMPORTANT)

Go provides heap via:
👉 `container/heap`

---

# 🔥 MIN HEAP IMPLEMENTATION

```go id="3zbc2s"
package main

import (
	"container/heap"
	"fmt"
)

// Define heap type
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] } // min heap
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}
```

---

# 🔥 MAX HEAP (just change Less)

```go id="vqdlk6"
func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
```

---

# 🚀 EXAMPLE USAGE

```go id="zfqk8c"
func main() {
	h := &MinHeap{3, 1, 5}
	heap.Init(h)

	heap.Push(h, 2)

	fmt.Println("Min:", (*h)[0])

	fmt.Println("Pop:", heap.Pop(h))
}
```

---

# 🧪 Output

```text
Min: 1
Pop: 1
```

---

# 🔥 MOST IMPORTANT PATTERNS

---

## 🟢 1. Top K Elements 🔥

👉 Keep heap of size k

---

### ✅ Example: Kth Largest

```go id="2yb8aq"
func findKthLargest(nums []int, k int) int {
	h := &MinHeap{}
	heap.Init(h)

	for _, v := range nums {
		heap.Push(h, v)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return (*h)[0]
}
```

---

## 🟢 2. Merge K Sorted Arrays 🔥

---

## 🟢 3. Median of Stream 🔥

👉 Two heaps:

* max heap (left)
* min heap (right)

---

## 🟢 4. Dijkstra Algorithm 🔥

---

# 🚀 FULL PROGRAM

```go id="9nvl8t"
package main

import (
	"container/heap"
	"fmt"
)

// ===== Heap =====
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	val := old[len(old)-1]
	*h = old[:len(old)-1]
	return val
}

// ===== Kth Largest =====
func findKthLargest(nums []int, k int) int {
	h := &MinHeap{}
	heap.Init(h)

	for _, v := range nums {
		heap.Push(h, v)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return (*h)[0]
}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}

	fmt.Println("Kth Largest:", findKthLargest(nums, 2))
}
```

---

# 🧪 Output

```text
Kth Largest: 5
```

---

# 🔥 INTERVIEW PATTERNS (VERY IMPORTANT)

## 🟢 Heap Size K Trick

👉 Use when:

* “top k”
* “kth smallest/largest”

---

## 🟢 Two Heaps Trick

👉 Median problems

---

## 🟢 Priority Queue

👉 Graphs (Dijkstra)

---

# 🎯 MOST ASKED HEAP PROBLEMS

* Kth Largest Element 🔥
* Top K Frequent Elements 🔥
* Merge K Sorted Lists 🔥
* Median of Data Stream 🔥
* Dijkstra Algorithm 🔥

---

# 🚀 Senior-Level Insight

👉 Heap = **partial sorting**

👉 Faster than sorting full array when:

```text
k << n
```

---

# 🚀 Next Step (FINAL CORE)

Now go to:

👉 **Graph (FINAL BOSS 🔥)**
👉 **Greedy**
👉 **DP (MOST IMPORTANT)**

---

Just say:

👉 **“graph start”**

You’re now entering **advanced interview level 🚀🔥**


Perfect — now you’re at the **FINAL BOSS of DSA: GRAPHS** 🔥
(If you master this → you’re interview-ready for top companies)

Let’s go **deep + pattern-based + Go implementation**

---

# 🌐 GRAPH — WHY IT EXISTS

👉 To represent **relationships between entities**

Examples:

* Network connections
* Social media
* Roads/maps
* Microservices communication (very relevant for you 🔥)

---

# 🧠 GRAPH TYPES

## 🔹 Directed vs Undirected

```text
A → B   (directed)
A — B   (undirected)
```

---

## 🔹 Weighted Graph

```text
A --5--> B
```

---

# 🧠 REPRESENTATION (VERY IMPORTANT)

---

## 🟢 1. Adjacency List (MOST USED)

```go id="xqkczj"
graph := make(map[int][]int)

graph[1] = []int{2, 3}
graph[2] = []int{4}
```

---

## 🟢 2. Adjacency Matrix

```text id="x4i81l"
  0 1 2
0 0 1 1
1 1 0 0
2 1 0 0
```

---

# 🔥 CORE GRAPH PATTERNS

---

## 🟢 1. DFS (Depth First Search)

👉 Uses recursion / stack

```go id="rbtbx0"
func dfs(node int, graph map[int][]int, visited map[int]bool) {
	visited[node] = true
	fmt.Print(node, " ")

	for _, nei := range graph[node] {
		if !visited[nei] {
			dfs(nei, graph, visited)
		}
	}
}
```

---

## 🟢 2. BFS (Breadth First Search) 🔥

👉 Uses queue

```go id="qr61yy"
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
```

---

# 🔥 IMPORTANT PROBLEMS

---

## 🟢 3. Number of Islands 🔥

👉 Matrix → graph

---

## 🟢 4. Cycle Detection

---

## 🟢 5. Topological Sort 🔥

👉 For DAG (dependency problems)

---

### ✅ Kahn’s Algorithm

```go id="7y4hfr"
func topoSort(n int, graph map[int][]int) []int {
	indegree := make([]int, n)

	for _, neighbors := range graph {
		for _, v := range neighbors {
			indegree[v]++
		}
	}

	queue := []int{}
	for i := 0; i < n; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	result := []int{}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		result = append(result, node)

		for _, nei := range graph[node] {
			indegree[nei]--
			if indegree[nei] == 0 {
				queue = append(queue, nei)
			}
		}
	}

	return result
}
```

---

## 🟢 6. Shortest Path 🔥

### Dijkstra (using heap)

---

# 🚀 FULL GO PROGRAM

```go id="6r9b4j"
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
```

---

# 🧪 Output

```text id="2e6d9w"
DFS: 0 1 2 3
BFS: 0 1 2 3
```

---

# 🔥 INTERVIEW PATTERNS (VERY IMPORTANT)

## 🟢 BFS → shortest path (unweighted)

---

## 🟢 DFS → connectivity, cycle

---

## 🟢 Topological Sort → dependencies

---

## 🟢 Union-Find → connected components

---

# 🎯 MOST ASKED GRAPH PROBLEMS

You MUST practice:

* Number of Islands 🔥
* Course Schedule 🔥
* Topological Sort 🔥
* Dijkstra 🔥
* Clone Graph

---

# 🚀 Senior-Level Insight

👉 Graph = **generalization of all data structures**

👉 Many problems:

```text id="fjr3dq"
Tree → special graph  
Matrix → implicit graph  
```

---

# 🚀 FINAL STEP (VERY IMPORTANT)

Now move to:

👉 **Greedy**
👉 **Backtracking**
👉 **Dynamic Programming (DP) 🔥🔥🔥**

---

Just say:

👉 **“dp deep dive”**

This is the **most important topic for cracking top interviews 🚀🔥**





Perfect — **Greedy is a super important pattern** 🔥
(very common in interviews because it tests your **intuition + proof thinking**)

Let’s go **deep + pattern-based + Go implementation**

---

# 🧠 GREEDY — WHY IT EXISTS

👉 Make the **best local choice**
👉 Hope it leads to **global optimal solution**

---

## 🎯 Key Idea

```text
At every step → choose the best option NOW
```

---

# ⚠️ IMPORTANT

👉 Greedy **does NOT always work**
👉 You must be able to **justify why it works**

---

# 🎯 WHERE GREEDY IS USED

* Scheduling problems
* Interval problems
* Resource allocation
* Optimization problems

---

# 🧠 CORE PATTERNS (VERY IMPORTANT)

---

## 🟢 1. Sort + Pick

👉 Most common greedy pattern

---

## 🟢 2. Earliest Finish Time 🔥

👉 Interval scheduling

---

## 🟢 3. Local Optimal Choice

👉 Example: Jump Game

---

## 🟢 4. Greedy + Heap

👉 Advanced problems

---

# 🔥 MOST IMPORTANT PROBLEMS

---

## 🟢 1. Activity Selection 🔥

### 🧠 Idea

Pick activity with **earliest end time**

---

### ✅ Go Code

```go
package main

import (
	"fmt"
	"sort"
)

type Activity struct {
	start, end int
}

func maxActivities(acts []Activity) int {
	sort.Slice(acts, func(i, j int) bool {
		return acts[i].end < acts[j].end
	})

	count := 1
	lastEnd := acts[0].end

	for i := 1; i < len(acts); i++ {
		if acts[i].start >= lastEnd {
			count++
			lastEnd = acts[i].end
		}
	}
	return count
}
```

---

## 🟢 2. Jump Game 🔥

👉 Can you reach end?

```go
func canJump(nums []int) bool {
	maxReach := 0

	for i := 0; i < len(nums); i++ {
		if i > maxReach {
			return false
		}
		if i+nums[i] > maxReach {
			maxReach = i + nums[i]
		}
	}
	return true
}
```

---

## 🟢 3. Minimum Coins

👉 Choose largest coin first

---

## 🟢 4. Gas Station 🔥

---

### ✅ Go Code

```go
func canCompleteCircuit(gas []int, cost []int) int {
	total, tank, start := 0, 0, 0

	for i := 0; i < len(gas); i++ {
		diff := gas[i] - cost[i]
		total += diff
		tank += diff

		if tank < 0 {
			start = i + 1
			tank = 0
		}
	}

	if total < 0 {
		return -1
	}
	return start
}
```

---

# 🚀 FULL PROGRAM

```go
package main

import (
	"fmt"
	"sort"
)

type Activity struct {
	start, end int
}

func maxActivities(acts []Activity) int {
	sort.Slice(acts, func(i, j int) bool {
		return acts[i].end < acts[j].end
	})

	count := 1
	lastEnd := acts[0].end

	for i := 1; i < len(acts); i++ {
		if acts[i].start >= lastEnd {
			count++
			lastEnd = acts[i].end
		}
	}
	return count
}

func canJump(nums []int) bool {
	maxReach := 0

	for i := 0; i < len(nums); i++ {
		if i > maxReach {
			return false
		}
		if i+nums[i] > maxReach {
			maxReach = i + nums[i]
		}
	}
	return true
}

func main() {

	acts := []Activity{
		{1, 3}, {2, 4}, {3, 5}, {0, 6}, {5, 7},
	}

	fmt.Println("Max Activities:", maxActivities(acts))

	fmt.Println("Can Jump:", canJump([]int{2, 3, 1, 1, 4}))
}
```

---

# 🧪 Output

```text
Max Activities: 3
Can Jump: true
```

---

# 🔥 INTERVIEW PATTERNS (VERY IMPORTANT)

## 🟢 Sort + Greedy

👉 80% greedy problems use sorting

---

## 🟢 Interval Problems

👉 Choose:

* earliest finish
* maximum overlap

---

## 🟢 Greedy + Prefix

👉 optimization problems

---

# 🎯 MOST ASKED GREEDY PROBLEMS

* Activity Selection 🔥
* Jump Game 🔥
* Gas Station 🔥
* Meeting Rooms 🔥
* Candy Distribution 🔥

---

# 🚀 Senior-Level Insight

👉 Always ask:

```text
Can I make a local decision that guarantees global optimal?
```

---

# ⚠️ GREEDY VS DP

| Greedy             | DP             |
| ------------------ | -------------- |
| Fast               | Slower         |
| Simple             | Complex        |
| Not always correct | Always correct |

---

# 🚀 FINAL STEP (MOST IMPORTANT)

Now go to:

👉 **Backtracking**
👉 **Dynamic Programming (DP) 🔥🔥🔥**

---

Just say:

👉 **“dp deep dive”**

This is the **most important topic for cracking top companies 🚀🔥**



Perfect — **Backtracking is the bridge between recursion → DP → hard problems** 🔥
(very commonly asked for **combinatorics + search problems**)

Let’s go **deep + pattern-based + Go implementation**

---

# 🧠 BACKTRACKING — WHY IT EXISTS

👉 Used when you need to:

```text
Try ALL possible choices
```

---

## 🎯 Core Idea

```text
Choose → Explore → Undo (Backtrack)
```

---

## 🧠 Why not brute force?

👉 Backtracking = **optimized brute force**

* Prunes invalid paths early
* Avoids unnecessary work

---

# 🎯 WHERE USED

* Permutations
* Combinations
* Subsets
* N-Queens 🔥
* Sudoku 🔥
* Word Search 🔥

---

# 🧠 CORE TEMPLATE (VERY IMPORTANT)

```go
func backtrack(path []int, choices []int) {
    // base case
    if condition {
        result = append(result, copy(path))
        return
    }

    for each choice {
        // choose
        path = append(path, choice)

        // explore
        backtrack(path, remaining)

        // undo
        path = path[:len(path)-1]
    }
}
```

👉 This template solves **80% backtracking problems**

---

# 🔥 MOST IMPORTANT PATTERNS

---

## 🟢 1. Subsets (Power Set)

### ✅ Go Code

```go id="k4y8qk"
func subsets(nums []int) [][]int {
	var res [][]int
	var path []int

	var backtrack func(int)
	backtrack = func(start int) {
		res = append(res, append([]int{}, path...))

		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtrack(0)
	return res
}
```

---

## 🟢 2. Permutations 🔥

```go id="hm1y7c"
func permute(nums []int) [][]int {
	var res [][]int
	var path []int
	used := make([]bool, len(nums))

	var backtrack func()
	backtrack = func() {
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}

			used[i] = true
			path = append(path, nums[i])

			backtrack()

			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtrack()
	return res
}
```

---

## 🟢 3. Combination Sum 🔥

```go id="nqz32x"
func combinationSum(nums []int, target int) [][]int {
	var res [][]int
	var path []int

	var backtrack func(int, int)
	backtrack = func(start int, sum int) {
		if sum == target {
			res = append(res, append([]int{}, path...))
			return
		}
		if sum > target {
			return
		}

		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(i, sum+nums[i])
			path = path[:len(path)-1]
		}
	}

	backtrack(0, 0)
	return res
}
```

---

# 🚀 FULL PROGRAM

```go id="q2g9bz"
package main

import "fmt"

func main() {

	fmt.Println("Subsets:", subsets([]int{1, 2, 3}))

	fmt.Println("Permutations:", permute([]int{1, 2, 3}))

	fmt.Println("Combination Sum:", combinationSum([]int{2, 3, 6, 7}, 7))
}

// ===== Subsets =====
func subsets(nums []int) [][]int {
	var res [][]int
	var path []int

	var backtrack func(int)
	backtrack = func(start int) {
		res = append(res, append([]int{}, path...))

		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtrack(0)
	return res
}

// ===== Permutations =====
func permute(nums []int) [][]int {
	var res [][]int
	var path []int
	used := make([]bool, len(nums))

	var backtrack func()
	backtrack = func() {
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}

			used[i] = true
			path = append(path, nums[i])

			backtrack()

			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtrack()
	return res
}

// ===== Combination Sum =====
func combinationSum(nums []int, target int) [][]int {
	var res [][]int
	var path []int

	var backtrack func(int, int)
	backtrack = func(start int, sum int) {
		if sum == target {
			res = append(res, append([]int{}, path...))
			return
		}
		if sum > target {
			return
		}

		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(i, sum+nums[i])
			path = path[:len(path)-1]
		}
	}

	backtrack(0, 0)
	return res
}
```

---

# 🧪 Output

```text
Subsets: [[], [1], [1 2], [1 2 3], [1 3], [2], [2 3], [3]]
Permutations: [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]
Combination Sum: [[2 2 3] [7]]
```

---

# 🔥 INTERVIEW PATTERNS (VERY IMPORTANT)

## 🟢 Choose → Explore → Undo

---

## 🟢 Use `start index`

👉 Avoid duplicates

---

## 🟢 Use `visited[]`

👉 For permutations

---

## 🟢 Pruning 🔥

```go
if sum > target { return }
```

👉 Huge optimization

---

# 🎯 MOST ASKED BACKTRACKING PROBLEMS

* Subsets 🔥
* Permutations 🔥
* Combination Sum 🔥
* N-Queens 🔥
* Sudoku Solver 🔥

---

# 🚀 Senior-Level Insight

👉 Backtracking = **DFS on decision tree**

---

# ⚠️ Complexity

👉 Usually:

```text
O(2^n) or O(n!)
```

---

# 🚀 FINAL STEP (MOST IMPORTANT)

Now comes:

👉 **Dynamic Programming (DP) 🔥🔥🔥**

---

Just say:

👉 **“dp deep dive”**

This is the **MOST IMPORTANT topic for cracking top interviews 🚀🔥**




Perfect — now you’ve reached the **most important topic in DSA: Dynamic Programming (DP)** 🔥🔥🔥

👉 If you master this → you can crack **top-tier interviews**

Let’s go **deep + structured + pattern-based (not confusing)**

---

# 🧠 DYNAMIC PROGRAMMING — WHY IT EXISTS

👉 To **avoid recomputation**

---

## 🔥 Core Idea

```text
Solve once → store → reuse
```

---

## 🧠 When DP is used

If problem has:

* **Overlapping subproblems**
* **Optimal substructure**

👉 Think DP

---

# 🧠 3 STEPS TO SOLVE DP (VERY IMPORTANT)

---

## 🟢 1. Define State

👉 What does `dp[i]` represent?

---

## 🟢 2. Transition

👉 How to compute from previous states?

---

## 🟢 3. Base Case

👉 Starting condition

---

# 🔥 TYPES OF DP

---

## 🟢 1. 1D DP

👉 Fibonacci, Climbing Stairs

---

## 🟢 2. 2D DP

👉 LCS, Grid problems

---

## 🟢 3. DP on Subsequences

👉 Knapsack, subset sum

---

## 🟢 4. DP on Trees

👉 Advanced

---

# 🚀 MOST IMPORTANT PATTERNS

---

## 🟢 1. Fibonacci (Basic DP)

```go id="gfkm3n"
func fib(n int) int {
	if n <= 1 {
		return n
	}

	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
```

---

## 🟢 2. Climbing Stairs 🔥

```go id="zph2lg"
func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
```

---

## 🟢 3. House Robber 🔥

```go id="qn2k8n"
func rob(nums []int) int {
	prev1, prev2 := 0, 0

	for _, num := range nums {
		temp := prev1
		if prev2+num > prev1 {
			prev1 = prev2 + num
		}
		prev2 = temp
	}
	return prev1
}
```

---

## 🟢 4. Longest Increasing Subsequence 🔥

```go id="o6jcb1"
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	maxLen := 1

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}
	return maxLen
}
```

---

## 🟢 5. 0/1 Knapsack 🔥🔥🔥

```go id="h9k4v8"
func knapsack(wt, val []int, W int) int {
	n := len(wt)
	dp := make([][]int, n+1)

	for i := range dp {
		dp[i] = make([]int, W+1)
	}

	for i := 1; i <= n; i++ {
		for w := 0; w <= W; w++ {
			if wt[i-1] <= w {
				include := val[i-1] + dp[i-1][w-wt[i-1]]
				exclude := dp[i-1][w]
				if include > exclude {
					dp[i][w] = include
				} else {
					dp[i][w] = exclude
				}
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}
	return dp[n][W]
}
```

---

# 🚀 FULL PROGRAM

```go id="r6c3xz"
package main

import "fmt"

func main() {

	fmt.Println("Fib:", fib(5))
	fmt.Println("Climb Stairs:", climbStairs(5))
	fmt.Println("House Robber:", rob([]int{2,7,9,3,1}))
	fmt.Println("LIS:", lengthOfLIS([]int{10,9,2,5,3,7,101,18}))

	wt := []int{1, 3, 4, 5}
	val := []int{1, 4, 5, 7}
	fmt.Println("Knapsack:", knapsack(wt, val, 7))
}

// ===== Fibonacci =====
func fib(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// ===== Climbing Stairs =====
func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// ===== House Robber =====
func rob(nums []int) int {
	prev1, prev2 := 0, 0

	for _, num := range nums {
		temp := prev1
		if prev2+num > prev1 {
			prev1 = prev2 + num
		}
		prev2 = temp
	}
	return prev1
}

// ===== LIS =====
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	maxLen := 1

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}
	return maxLen
}

// ===== Knapsack =====
func knapsack(wt, val []int, W int) int {
	n := len(wt)
	dp := make([][]int, n+1)

	for i := range dp {
		dp[i] = make([]int, W+1)
	}

	for i := 1; i <= n; i++ {
		for w := 0; w <= W; w++ {
			if wt[i-1] <= w {
				include := val[i-1] + dp[i-1][w-wt[i-1]]
				exclude := dp[i-1][w]
				if include > exclude {
					dp[i][w] = include
				} else {
					dp[i][w] = exclude
				}
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}
	return dp[n][W]
}
```

---

# 🧪 Output

```text
Fib: 5
Climb Stairs: 8
House Robber: 12
LIS: 4
Knapsack: 9
```

---

# 🔥 INTERVIEW PATTERNS (VERY IMPORTANT)

## 🟢 1. “Count ways” → DP

---

## 🟢 2. “Min / Max” → DP

---

## 🟢 3. “Subset / Combination” → DP

---

## 🟢 4. “Take / Not Take” → DP

---

# 🧠 GREEDY vs DP (VERY IMPORTANT)

| Greedy             | DP                  |
| ------------------ | ------------------- |
| Fast               | Slower              |
| Not always correct | Always correct      |
| Local decision     | Global optimization |

---

# 🚀 SENIOR-LEVEL INSIGHT

👉 DP = **memoized recursion OR bottom-up table**

👉 Always think:

```text
What is my state?
```

---

# 🚀 FINAL ROADMAP (YOU COMPLETED CORE DSA 🎉)

You now know:

* Arrays
* Strings
* Hashing
* Stack / Queue / Deque
* Linked List
* Tree / BST
* Heap
* Graph
* Greedy
* Backtracking
* DP

👉 This is **full DSA coverage**

---

# 🚀 NEXT LEVEL (IMPORTANT)

Now you should:

👉 Solve **pattern-wise problems**
👉 Practice **mock interviews**
👉 Do **company-specific questions (PubMatic level)**