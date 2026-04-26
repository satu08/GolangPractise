package main

import "fmt"

func main() {

	fmt.Println("Fib:", fib(5))
	fmt.Println("Climb Stairs:", climbStairs(5))
	fmt.Println("House Robber:", rob([]int{2, 7, 9, 3, 1}))
	fmt.Println("LIS:", lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))

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
