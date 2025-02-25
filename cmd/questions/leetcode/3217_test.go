package main

import (
	"reflect"
	"testing"

	libs "github.com/meetaayush/dsa-in-go/cmd/libs/linkedList"
)

func Test3217(t *testing.T) {
	type testCase struct {
		list   []int
		nums   []int
		result []int
	}

	cases := []testCase{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3}, []int{4, 5}},
		{[]int{1, 2, 1, 2, 1, 2}, []int{1}, []int{2, 2, 2}},
		{[]int{1, 2, 3, 4}, []int{5}, []int{1, 2, 3, 4}},
	}
	for _, val := range cases {
		list1 := &libs.LinkedList{}
		for _, numVal1 := range val.list {
			tempNode := &libs.Node{}
			tempNode.Val = numVal1
			list1.Append(tempNode)
		}
		res := modifiedList(list1.Head, val.nums)
		expectedRes := res.ParseToArray()

		if !reflect.DeepEqual(expectedRes, val.result) {
			t.Errorf("expected %v, got %v", val.result, expectedRes)
		}
	}
}
