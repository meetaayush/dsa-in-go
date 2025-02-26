package main

import "testing"

func Test1749(t *testing.T) {
	type testCase struct {
		nums   []int
		result int
	}

	cases := []testCase{
		{[]int{1, -3, 2, 3, -4}, 5},
		{[]int{2, -5, 1, -4, 3, -2}, 8},
	}

	for _, item := range cases {
		res := maxAbsoluteSum(item.nums)
		if res != item.result {
			t.Errorf("expected %v, got %v", item.result, res)
		}
	}
}
