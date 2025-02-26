package main

import "math"

func findMin(arr []int) int {
    left, right, ans := 0, len(arr) - 1, 0
    target := arr[right]
    for left <= right {
        mid := left + int(math.Floor(float64(right - left) / 2))
        guess := arr[mid]
        if guess <= target {
            ans = mid
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return arr[ans]
}
