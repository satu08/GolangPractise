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
