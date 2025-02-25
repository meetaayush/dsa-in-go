package main

import (
	"reflect"
	"testing"

	libs "github.com/meetaayush/dsa-in-go/cmd/libs/linkedList"
)

func TestSortLL012(t *testing.T) {
	type testCase struct {
		list   []int
		result []int
	}

	cases := []testCase{
		{[]int{1, 0, 2, 0, 1}, []int{0, 0, 1, 1, 2}},
		{[]int{1, 1, 1, 0}, []int{0, 1, 1, 1}},
		{[]int{2, 2, 1, 2}, []int{1, 2, 2, 2}},
	}

	for _, val := range cases {
		list := &libs.LinkedList{}
		for _, item := range val.list {
			node := &libs.Node{}
			node.Val = item

			list.Append(node)
		}

		res := sortList(list.Head)
		arr := res.ParseToArray()

		if !reflect.DeepEqual(arr, val.result) {
			t.Errorf("expected %v, got %v", val.result, arr)
		}
	}
}
