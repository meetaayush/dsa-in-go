package main

import libs "github.com/meetaayush/dsa-in-go/cmd/libs/linkedList"

func sortList(head *libs.Node) *libs.Node {
	if head == nil {
		return head
	}

	dummyOnes := &libs.Node{}
	dummyTwos := &libs.Node{}
	dummyZeros := &libs.Node{}

	zero := dummyZeros
	one := dummyOnes
	two := dummyTwos

	curr := head
	for curr != nil {
		if curr.Val == 0 {
			zero.Next = curr
			zero = zero.Next
		} else if curr.Val == 1 {
			one.Next = curr
			one = one.Next
		} else {
			two.Next = curr
			two = two.Next
		}

		curr = curr.Next
	}

	if dummyOnes.Next != nil {
		zero.Next = dummyOnes.Next
	} else {
		zero.Next = dummyTwos.Next
	}

	one.Next = dummyTwos.Next
	two.Next = nil

	return dummyZeros.Next
}
