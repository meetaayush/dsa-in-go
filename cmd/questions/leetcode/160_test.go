package main

import (
	"reflect"
	"testing"

	libs "github.com/meetaayush/dsa-in-go/cmd/libs/linkedList"
)

func Test160(t *testing.T) {
	type testCase struct {
		list1            []int
		list2            []int
		intersectionIdx1 int
		intersectionIdx2 int
		result           int
	}

	cases := []testCase{
		{[]int{4, 1, 8, 4, 5}, []int{5, 6, 1, 8, 4, 5}, 2, 3, 8},
	}
	for _, val := range cases {
		list1, list2 := &libs.LinkedList{}, &libs.LinkedList{}
		var interSectionNode1 *libs.Node
		for idx1, numVal1 := range val.list1 {
			tempNode := &libs.Node{}
			tempNode.Val = numVal1

			if idx1 == val.intersectionIdx1 {
				interSectionNode1 = tempNode
			}

			list1.Append(tempNode)
		}
		for idx2, numVal2 := range val.list2 {
			if idx2 == val.intersectionIdx2 {
				list2.Append(interSectionNode1)
				break
			}
			tempNode := &libs.Node{}
			tempNode.Val = numVal2
			list2.Append(tempNode)
		}

		res := getIntersectionNode(list1.Head, list2.Head)

		if res == nil && val.result != -1 {
			t.Errorf("expected %v, got %v", val.result, nil)
		}

		if !reflect.DeepEqual(res.Val, val.result) {
			t.Errorf("expected %v, got %v", val.result, res.Val)
		}
	}
}
