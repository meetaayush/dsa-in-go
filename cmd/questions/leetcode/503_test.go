package main

import (
	"reflect"
	"testing"
)

func Test503(t *testing.T) {
	type testCase struct {
		nums     []int
		expected []int
	}
	cases := []testCase{
		{[]int{1, 2, 1}, []int{2, -1, 2}},
		{[]int{1, 2, 3, 4, 3}, []int{2, 3, 4, -1, 4}},
		{[]int{-1}, []int{-1}},
	}

	for _, val := range cases {
		res := nextGreaterElements(val.nums)
		if !reflect.DeepEqual(res, val.expected) {
			t.Errorf("expected %v, got %v", val.expected, res)
		}
	}
}
