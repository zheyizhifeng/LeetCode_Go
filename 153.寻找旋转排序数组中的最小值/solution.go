// @algorithm @lc id=153 lang=golang 
// @title find-minimum-in-rotated-sorted-array


package main
// @test([3,4,5,1,2])=1
// @test([4,5,6,7,0,1,2])=0
// @test([11,13,15,17])=11
func findMin(nums []int) int {
	length := len(nums)
	if nums[0] <= nums[length-1] { return nums[0] }
	left, right := 0, length - 1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] >= nums[0] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}