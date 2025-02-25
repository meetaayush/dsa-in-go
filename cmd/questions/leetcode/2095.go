package main

import libs "github.com/meetaayush/dsa-in-go/cmd/libs/linkedList"

func deleteMiddle(head *libs.Node) *libs.Node {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return nil
	}

	slow, fast := head, head.Next.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// slow is at mid - 1
	slow.Next = slow.Next.Next

	return head
}
