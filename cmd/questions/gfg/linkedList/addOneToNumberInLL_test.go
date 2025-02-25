package main

import (
	"reflect"
	"testing"

	libs "github.com/meetaayush/dsa-in-go/cmd/libs/linkedList"
)

func TestAddOneToNumberInLL(t *testing.T) {
	type testCase struct {
		num    []int
		result []int
	}

	cases := []testCase{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 6}},
		{[]int{2, 1, 3, 5, 6, 9, 9}, []int{2, 1, 3, 5, 7, 0, 0}},
		{[]int{9}, []int{1, 0}},
		{[]int{1}, []int{2}},
	}
	for _, val := range cases {
		list1 := &libs.LinkedList{}
		for _, numVal1 := range val.num {
			tempNode := &libs.Node{}
			tempNode.Val = numVal1
			list1.Append(tempNode)
		}
		res := addOneToNumberInLL(list1.Head)
		expectedRes := res.ParseToArray()

		if !reflect.DeepEqual(expectedRes, val.result) {
			t.Errorf("expected %v, got %v", val.result, expectedRes)
		}
	}
}
