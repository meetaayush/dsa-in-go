package main

import "math"

func peakIndexInMountainArray(arr []int) int {
    left, right, ans := 0, len(arr) - 1, 0
    for left <= right {
        mid := left + int(math.Floor(float64(right - left) / 2))

        if arr[mid] >= arr[mid + 1] {
            ans = mid
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return ans
}
