// @algorithm @lc id=882 lang=golang
// @title peak-index-in-a-mountain-array

package main

// @test([0,1,0])=1
// @test([0,2,1,0])=1
// @test([0,10,5,2])=1
func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)-1
	for left < right {
		mid := left + (right-left+1)/2
		if arr[mid] > arr[mid-1] {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}
