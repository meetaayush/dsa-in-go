package main

import (
	"reflect"
	"testing"

	libs "github.com/meetaayush/dsa-in-go/cmd/libs/linkedList"
)

func TestRemoveNthFromEndOfLL(t *testing.T) {
	type testCase struct {
		list   []int
		k      int
		result []int
	}

	cases := []testCase{
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
		{[]int{1}, 1, []int{}},
		{[]int{1, 2}, 1, []int{1}},
		{[]int{1, 2}, 2, []int{2}},
	}

	for _, val := range cases {
		list := &libs.LinkedList{}
		for _, num := range val.list {
			node := &libs.Node{}
			node.Val = num

			list.Append(node)
		}

		res := removeNthFromEnd(list.Head, val.k)
		expectedRes := res.ParseToArray()

		if !reflect.DeepEqual(expectedRes, val.result) {
			t.Errorf("expected %v, got %v", val.result, expectedRes)
		}
	}
}
