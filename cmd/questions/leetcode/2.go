// https://leetcode.com/problems/add-two-numbers/description/

package main

import (
	"math"

	libs "github.com/meetaayush/dsa-in-go/cmd/libs/linkedList"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers_2(l1 *libs.Node, l2 *libs.Node) *libs.Node {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	dummyNode := &libs.Node{}
	curr := dummyNode
	l1Ptr, l2Ptr, carry := l1, l2, 0
	for l1Ptr != nil && l2Ptr != nil {
		num1, num2 := l1Ptr.Val, l2Ptr.Val
		res := num1 + num2 + carry

		unit := res % 10
		tens := int(math.Floor(float64(res / 10)))

		node := &libs.Node{}
		node.Val = unit

		curr.Next = node
		carry = tens

		curr = curr.Next
		l1Ptr = l1Ptr.Next
		l2Ptr = l2Ptr.Next
	}

	for l1Ptr == nil && l2Ptr != nil {
		res := l2Ptr.Val + carry

		unit := res % 10
		tens := int(math.Floor(float64(res / 10)))

		node := &libs.Node{}
		node.Val = unit

		curr.Next = node
		carry = tens

		curr = curr.Next
		l2Ptr = l2Ptr.Next
	}

	for l2Ptr == nil && l1Ptr != nil {
		res := l1Ptr.Val + carry

		unit := res % 10
		tens := int(math.Floor(float64(res / 10)))

		node := &libs.Node{}
		node.Val = unit

		curr.Next = node
		carry = tens

		curr = curr.Next
		l1Ptr = l1Ptr.Next
	}
	if carry != 0 {
		node := &libs.Node{}
		node.Val = carry

		curr.Next = node
		curr = curr.Next
	}
	return dummyNode.Next
}
