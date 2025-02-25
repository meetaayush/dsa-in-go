// https://www.geeksforgeeks.org/add-1-number-represented-linked-list/

package main

import (
	"math"

	libs "github.com/meetaayush/dsa-in-go/cmd/libs/linkedList"
)

func addOneToNumberInLL(head *libs.Node) *libs.Node {
	if head == nil {
		return head
	}
	reversedHead := head.Reverse()
	curr, carry := reversedHead, 1

	for curr != nil {
		res := curr.Val + carry
		unit := res % 10
		tens := math.Floor(float64(res / 10))

		carry = int(tens)
		curr.Val = unit

		if carry == 0 {
			break
		}

		if curr.Next == nil && carry != 0 {
			node := &libs.Node{}
			node.Val = carry

			curr.Next = node
			break
		}

		curr = curr.Next
	}

	return reversedHead.Reverse()
}
