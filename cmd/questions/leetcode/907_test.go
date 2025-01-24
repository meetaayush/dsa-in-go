package main

import "testing"

func Test907(t *testing.T) {
	type testCase struct {
		nums     []int
		expected int
	}
	cases := []testCase{
		// {[]int{3, 1, 2, 4}, 17},
		// {[]int{3, 1, 2, 5}, 18},
		{[]int{11, 81, 94, 43, 3}, 444},
		// {[]int{2, 3, 1}, 10},
	}
	for _, val := range cases {
		got := sumSubarrayMins(val.nums)
		if got != val.expected {
			t.Errorf("expected %v, got %v", val.expected, got)
		}
	}
}
