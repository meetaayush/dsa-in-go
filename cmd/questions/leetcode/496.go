// https://leetcode.com/problems/next-greater-element-i/description/
package main

import (
	"fmt"

	libs "github.com/meetaayush/dsa-in-go/cmd/libs/stack"
)

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := nge(nums2)
	ans := make([]int, len(nums1))

	for idx1, item1 := range nums1 {
		ans[idx1] = m[item1]
	}

	return ans
}

func nge(nums []int) map[int]int {
	m := make(map[int]int)
	n := len(nums)
	st := libs.NewStack(n)
	for i := n - 1; i >= 0; i-- {
		// pop all elements from stack which are smaller or equal current element
		for st.IsEmpty() == false {
			topElementOnStack, _ := st.Peek()
			if topElementOnStack <= nums[i] {
				st.Pop()
			} else {
				break
			}
		}

		topElementOnStack, _ := st.Peek()
		if st.IsEmpty() {
			m[nums[i]] = -1
		} else {
			m[nums[i]] = topElementOnStack
		}
		st.Push(nums[i])
	}

	return m
}

func main() {
	arr1 := []int{4, 1, 2}
	arr2 := []int{1, 3, 4, 2}

	res := nextGreaterElement(arr1, arr2)
	fmt.Println(res)
}
