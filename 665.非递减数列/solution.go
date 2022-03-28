// @algorithm @lc id=665 lang=golang
// @title non-decreasing-array

package main

// @test([4,2,3])=true
// @test([4,2,1])=false
func checkPossibility(nums []int) bool {
	length := len(nums)
	changeCount := 0
	for i := 0; i < length-1; i++ {
		if nums[i] > nums[i+1] {
			// * 两种改变元素方式
			if (i >= 1 && nums[i+1] < nums[i-1]) && (i <= length-3 && nums[i] > nums[i+2]) {
				return false
			}
			changeCount++
			if changeCount > 1 {
				return false
			}
		}
	}
	return true
}
