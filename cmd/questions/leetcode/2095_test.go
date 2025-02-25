package main

import (
	"reflect"
	"testing"

	libs "github.com/meetaayush/dsa-in-go/cmd/libs/linkedList"
)

func Test2095(t *testing.T) {
	type testCase struct {
		nums     []int
		expected []int
	}
	cases := []testCase{
		{[]int{1, 2, 1}, []int{1, 1}},
		{[]int{1, 2, 3, 4, 3}, []int{1, 2, 4, 3}},
		{[]int{-1}, []int{}},
	}

	for _, val := range cases {
		list := &libs.LinkedList{}
		for _, item := range val.nums {
			node := new(libs.Node)
			node.Val = item

			list.Append(node)
		}
		res := deleteMiddle(list.Head)
		expectedRes := res.ParseToArray()
		if !reflect.DeepEqual(expectedRes, val.expected) {
			t.Errorf("expected %v, got %v", val.expected, expectedRes)
		}
	}
}
