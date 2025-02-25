package main

import libs "github.com/meetaayush/dsa-in-go/cmd/libs/linkedList"

func getIntersectionNode(headA, headB *libs.Node) *libs.Node {
	d1, d2 := headA, headB

	for d1 != d2 {
		if d1 == nil {
			d1 = headB
		} else {
			d1 = d1.Next
		}

		if d2 == nil {
			d2 = headA
		} else {
			d2 = d2.Next
		}
	}

	return d1
}
