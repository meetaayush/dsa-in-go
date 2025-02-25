package main

import libs "github.com/meetaayush/dsa-in-go/cmd/libs/linkedList"

func hasCycle(head *libs.Node) bool {
	if head == nil {
		return false
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}
