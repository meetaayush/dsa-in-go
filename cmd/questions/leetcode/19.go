// https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/

package main

import libs "github.com/meetaayush/dsa-in-go/cmd/libs/linkedList"

func removeNthFromEnd(head *libs.Node, n int) *libs.Node {
	if head == nil {
		return head
	}
	if head.Next == nil && n == 1 {
		return nil
	}
	slowPtr, fastPtr := head, head
	for i := 0; i < n; i++ {
		fastPtr = fastPtr.Next
	}

	// if fastPtr == nil then head needs to be removed
	if fastPtr == nil {
		temp := slowPtr.Next
		slowPtr.Next = nil
		head = temp

		return head
	}

	for fastPtr != nil && fastPtr.Next != nil {
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next
	}

	// delete next node of slowPtr
	temp := slowPtr.Next.Next
	slowPtr.Next.Next = nil
	slowPtr.Next = temp

	return head
}
