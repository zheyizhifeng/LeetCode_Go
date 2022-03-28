// @algorithm @lc id=81 lang=golang 
// @title search-in-rotated-sorted-array-ii


package main
// @test([2,5,6,0,0,1,2],0)=true
// @test([2,5,6,0,0,1,2],3)=false
func search(nums []int, target int) bool {
	length := len(nums)
	left, right := 0, length - 1
	for left < right {
		mid := left + (right - left) / 2
		if target == nums[mid] { return true }
		if nums[mid] == nums[left] {
			left++
			continue
		}
		if (nums[mid] > nums[left]) {
			if (target >= nums[left] && target < nums[mid]) {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if (target > nums[mid] && target <= nums[right]) {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return nums[left] == target
}