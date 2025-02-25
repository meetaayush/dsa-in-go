package main

import (
	"reflect"
	"testing"
)

func Test496(t *testing.T) {
	type TestCase struct {
		nums1    []int
		nums2    []int
		expected []int
	}
	cases := []TestCase{
		{[]int{4, 1, 2}, []int{1, 3, 4, 2}, []int{-1, 3, -1}},
		{[]int{2, 4}, []int{1, 2, 3, 4}, []int{3, -1}},
	}
	for _, val := range cases {
		res := nextGreaterElement(val.nums1, val.nums2)
		if !reflect.DeepEqual(res, val.expected) {
			t.Errorf("expected %v, got %v", val.expected, res)
		}
	}

}
