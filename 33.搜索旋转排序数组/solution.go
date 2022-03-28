// @algorithm @lc id=33 lang=golang
// @title search-in-rotated-sorted-array

package main

// @test([4,5,6,7,0,1,2],0)=4
// @test([4,5,6,7,0,1,2],3)=-1
// @test([1],0)=-1
func findPeakIndex(arr []int) int {
	left, right := 0, len(arr) - 1
	for left < right {
		mid := left + (right - left + 1)/2
		if arr[mid] >= arr[0] {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}
func binarySearch(arr []int, left, right, target int) int {
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if arr[left] == target {
		return left
	}
	return -1
}
func search(nums []int, target int) int {
	length := len(nums)
	if nums[0] <= nums[length-1] {
		return binarySearch(nums, 0, length-1, target)
	}
	peakIndex := findPeakIndex(nums)
	if target >= nums[0] && target <= nums[peakIndex] {
		return binarySearch(nums, 0, peakIndex, target)
	} else if target >= nums[peakIndex+1] && target <= nums[length-1] {
		return binarySearch(nums, peakIndex+1, length-1, target)
	}
	return -1
}
