package main

import libs "github.com/meetaayush/dsa-in-go/cmd/libs/linkedList"

func isPalindromeList(head *libs.Node) bool {
	if head == nil {
		return false
	}
	if head.Next == nil {
		return true
	}
	mid := getMid(head)
	newHead, curr := mid.Reverse(), head

	for newHead != nil {
		if curr.Val != newHead.Val {
			return false
		}
		curr = curr.Next
		newHead = newHead.Next
	}
	return true
}

func getMid(head *libs.Node) *libs.Node {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow.Next
}
