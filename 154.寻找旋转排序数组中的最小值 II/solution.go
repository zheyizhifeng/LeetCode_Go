// @algorithm @lc id=154 lang=golang 
// @title find-minimum-in-rotated-sorted-array-ii


package main
// @test([1,3,5])=1
// @test([2,2,2,0,1])=0
func findMin(nums []int) int {
	length := len(nums)
	left, right := 0, length - 1;
	for left < right {
		mid := left + (right - left)/2
		if nums[mid] == nums[right] {
			right--
		} else if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}