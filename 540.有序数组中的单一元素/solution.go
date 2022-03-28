// @algorithm @lc id=540 lang=golang
// @title single-element-in-a-sorted-array

package main

// @test([1,1,2,3,3,4,4,8,8])=2
// @test([3,3,7,7,10,11,11])=10
func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 { return nums[0] }
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] != nums[mid-1] && nums[mid] != nums[mid+1] {
			return nums[mid]
		}
		if nums[mid] == nums[mid-1] {
			if (mid & 1) == 1 {
				// * mid 前面奇数长度
				left = mid + 1
			} else {
				right = mid - 2
			}
		} else {
			if (mid & 1) == 1 {
				right = mid - 1
			} else {
				left = mid + 2
			}
		}
	}
	return nums[left]
}
