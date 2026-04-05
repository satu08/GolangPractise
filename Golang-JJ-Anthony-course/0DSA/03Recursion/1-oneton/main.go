package main

import "fmt"

func main() {
	nto1(5)
	fmt.Println("")
	oneton(5)
	fmt.Println("")
	onetontotailrecursive(5, 1)
}

// when caller has nothing to do after child call is called tail recursion
func nto1(n int) {
	if n == 0 {
		return
	}
	fmt.Printf("%v\t", n)
	nto1(n - 1)
}

// this takes auxiliary space O(n)
func oneton(n int) {
	if n == 0 {
		return
	}
	oneton(n - 1)
	fmt.Printf("%v\t", n)
}

// equivalent tail recursive for oneton
func onetontotailrecursive(n, k int) {
	if n == 0 {
		return
	}
	fmt.Printf("%v\t", k)
	onetontotailrecursive(n-1, k+1)
}
