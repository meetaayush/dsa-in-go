package main

import (
	"reflect"
	"testing"

	libs "github.com/meetaayush/dsa-in-go/cmd/libs/linkedList"
)

func Test234(t *testing.T) {
	type testCase struct {
		num    []int
		result bool
	}

	cases := []testCase{
		{[]int{1, 2, 2, 1}, true},
		{[]int{2, 1}, false},
		{[]int{9}, true},
		{[]int{}, false},
	}
	for _, val := range cases {
		list1 := &libs.LinkedList{}
		for _, numVal1 := range val.num {
			tempNode := &libs.Node{}
			tempNode.Val = numVal1
			list1.Append(tempNode)
		}
		res := isPalindromeList(list1.Head)

		if !reflect.DeepEqual(res, val.result) {
			t.Errorf("expected %v, got %v", val.result, res)
		}
	}
}
