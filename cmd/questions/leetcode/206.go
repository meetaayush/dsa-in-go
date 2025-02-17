package main

import libs "github.com/meetaayush/dsa-in-go/cmd/libs/linkedList"

func reverseList(head *libs.Node) *libs.Node {
	if head == nil {
		return head
	}

	prev, curr := &libs.Node{}, head
	prev = nil

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
