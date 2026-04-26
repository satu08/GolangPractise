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
