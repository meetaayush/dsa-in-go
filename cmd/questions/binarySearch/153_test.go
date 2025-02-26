package main

import "testing"

func Test153(t *testing.T) {
	type testCase struct {
		nums   []int
		result int
	}

	cases := []testCase{
		{[]int{3, 4, 5, 1, 2}, 1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
	}

	for _, item := range cases {
		res := findMin(item.nums)
		if res != item.result {
			t.Errorf("expected %v, got %v", item.result, res)
		}
	}
}
