package main

import (
	"reflect"
	"testing"

	libs "github.com/meetaayush/dsa-in-go/cmd/libs/linkedList"
)

func Test2(t *testing.T) {
	type testCase struct {
		num1   []int
		num2   []int
		result []int
	}

	cases := []testCase{
		{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{[]int{0}, []int{0}, []int{0}},
		{[]int{9, 9, 9, 9, 9}, []int{}, []int{9, 9, 9, 9, 9}},
	}

	for _, val := range cases {
		list1 := &libs.LinkedList{}
		for _, numVal1 := range val.num1 {
			tempNode := &libs.Node{}
			tempNode.Val = numVal1
			list1.Append(tempNode)
		}

		list2 := &libs.LinkedList{}
		for _, numVal2 := range val.num2 {
			tempNode := &libs.Node{}
			tempNode.Val = numVal2
			list2.Append(tempNode)
		}

		res := addTwoNumbers_2(list1.Head, list2.Head)
		if !reflect.DeepEqual(res.ParseToArray(), val.result) {
			t.Errorf("expected %v, Got %v", val.result, res.ParseToArray())
		}
	}
}
