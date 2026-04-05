package main

import "testing"

func TestAdd(t *testing.T) {
	type test struct {
		data []int
		ans  int
	}
	// table test
	tests := []test{
		test{[]int{5, 6, 3}, 14},
		test{[]int{9, -1, 0}, 8},
		test{[]int{11, 23, -33}, 1},
		test{[]int{1, -1, 0}, 0},
	}

	for _, test := range tests {
		total := add(test.data...)
		if total != test.ans {
			t.Errorf("sum was incorrect, got %d want %d", test.ans, total)
		}
	}

}
