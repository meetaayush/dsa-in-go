// https://leetcode.com/problems/delete-nodes-from-linked-list-present-in-array/description
package main

import libs "github.com/meetaayush/dsa-in-go/cmd/libs/linkedList"

func modifiedList(head *libs.Node, nums []int) *libs.Node {
	if head == nil {
		return head
	}
	numsBoolArray := [1e5]bool{}
	for _, val := range nums {
		numsBoolArray[val] = true
	}

	dummyNode := &libs.Node{}
	dummyNode.Next = head

	curr := dummyNode

	for curr.Next != nil {
		if numsBoolArray[curr.Next.Val] {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

	return dummyNode.Next
}
