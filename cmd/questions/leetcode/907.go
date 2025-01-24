package main

import (
	libs "github.com/meetaayush/dsa-in-go/cmd/libs/stack"
)

// sumSubarrayMins uses the monotonic stack to find
// the sum of subarray Mins
func sumSubarrayMins(arr []int) int {
	arrNse := nse(arr)
	arrPse := pse(arr)
	mod := int(1e9 + 7)
	n := len(arr)

	total := 0

	for i := 0; i < n; i++ {
		left := i - arrPse[i]
		right := arrNse[i] - i

		freq := left * right * 1
		val := (freq * arr[i]) % mod

		total = (total + val) % mod

		// fmt.Printf("left = %v\nRight = %v\nTotal = %v\n\n", left, right, total)
	}

	return total
}

// sumSubarrayMinsBrute uses the brute force approach to find
// the sum of subArray minimums
func sumSubarrayMinsBrute(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		minElem := arr[i]
		for j := i; j < len(arr); j++ {
			minElem = min(minElem, arr[j])
			sum = sum + minElem
		}
	}

	return sum
}

// min will find the minimum of minimum of num1 and num2
func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	} else {
		return num2
	}
}

// nse finds the next smaller element for
// each element of the input array
func nse(arr []int) []int {
	n := len(arr)
	st := libs.NewStack(n)
	ans := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		for !st.IsEmpty() {
			topElem, _ := st.Peek()
			if arr[topElem] >= arr[i] {
				st.Pop()
			} else {
				break
			}
		}

		if st.IsEmpty() {
			ans[i] = n
		} else {
			topElem, _ := st.Peek()
			ans[i] = topElem
		}

		st.Push(i)
	}

	return ans
}

// pse finds the previous smaller element for
// each element of the input array
func pse(arr []int) []int {
	n := len(arr)
	st := libs.NewStack(n)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		for !st.IsEmpty() {
			topElem, _ := st.Peek()
			if arr[topElem] > arr[i] {
				st.Pop()
			} else {
				break
			}
		}

		if st.IsEmpty() {
			ans[i] = -1
		} else {
			topElem, _ := st.Peek()
			ans[i] = topElem
		}

		st.Push(i)
	}

	return ans
}
