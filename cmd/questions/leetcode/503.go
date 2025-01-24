// https://leetcode.com/problems/next-greater-element-ii/

package main

import (
	libs "github.com/meetaayush/dsa-in-go/cmd/libs/stack"
)

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	st := libs.NewStack(n)
	for i := range ans {
		ans[i] = -1
	}
	for i := (2 * n) - 1; i >= 0; i-- {
		for !st.IsEmpty() {
			topElement, _ := st.Peek()
			if topElement <= nums[i%n] {
				st.Pop()
			} else {
				break
			}
		}

		if i < n {
			if st.IsEmpty() {
				ans[i] = -1
			} else {
				topElement, _ := st.Peek()
				ans[i] = topElement
			}
		}
		st.Push(nums[i%n])
	}
	return ans
}
