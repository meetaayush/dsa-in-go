package main

import libs "github.com/meetaayush/dsa-in-go/cmd/libs/linkedList"

func oddEvenList(head *libs.Node) *libs.Node {
	if head == nil {
		return nil
	}
	odd := head
	even := odd.Next
	firstEven := head.Next

	for even != nil && even.Next != nil {
		odd.Next = odd.Next.Next
		even.Next = even.Next.Next
		odd = odd.Next
		even = even.Next
	}

	odd.Next = firstEven

	return head
}
