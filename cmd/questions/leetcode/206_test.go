package main

import (
	"reflect"
	"testing"

	libs "github.com/meetaayush/dsa-in-go/cmd/libs/linkedList"
)

func TestReverseLL(t *testing.T) {
	type testCase struct {
		list   []int
		result []int
	}

	cases := []testCase{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{1}, []int{1}},
		{[]int{6, 8}, []int{8, 6}},
		{[]int{}, []int{}},
	}

	for _, val := range cases {
		list := &libs.LinkedList{}
		for _, num := range val.list {
			node := &libs.Node{}
			node.Val = num

			list.Append(node)
		}

		res := reverseList(list.Head)
		expectedRes := res.ParseToArray()

		if !reflect.DeepEqual(expectedRes, val.result) {
			t.Errorf("expected %v, got %v", val.result, expectedRes)
		}
	}
}
