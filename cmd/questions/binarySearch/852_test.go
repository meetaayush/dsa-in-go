package main

import "testing"

func Test852(t *testing.T) {
	type testCase struct {
		nums   []int
		result int
	}

	cases := []testCase{
		{[]int{0, 1, 0}, 1},
		{[]int{0, 2, 1, 0}, 1},
		{[]int{0, 10, 5, 2}, 1},
	}

	for _, item := range cases {
		res := peakIndexInMountainArray(item.nums)
		if res != item.result {
			t.Errorf("expected %v, got %v", item.result, res)
		}
	}
}
