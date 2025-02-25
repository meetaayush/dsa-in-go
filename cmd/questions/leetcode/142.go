// https://leetcode.com/problems/linked-list-cycle-ii/

package main

import libs "github.com/meetaayush/dsa-in-go/cmd/libs/linkedList"

func detectCycle(head *libs.Node) *libs.Node {
	if head == nil {
		return nil
	}

	slow, fast, isLoop := head, head, false
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			isLoop = true
			break
		}
	}

	if !isLoop {
		return nil
	}

	slow = head

	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}
