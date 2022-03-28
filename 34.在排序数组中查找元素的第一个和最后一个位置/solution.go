// @algorithm @lc id=34 lang=golang 
// @title find-first-and-last-position-of-element-in-sorted-array


package main
// @test([5,7,7,8,8,10],8)=[3,4]
// @test([5,7,7,8,8,10],6)=[-1,-1]
// @test([],0)=[-1,-1]
func findFirstPos(nums []int, target int) int {
	l, r := 0, len(nums) - 1
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] == target {
			r = mid
		} else {
			r = mid - 1
		}
	}
	if nums[l] != target {
		return -1
	}
	return l
}

func findLastPos(nums []int, target int) int {
	l, r := 0, len(nums) - 1
	for l < r {
		mid := (l + r + 1) >> 1
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] == target {
			l = mid // * 左边界向下取整了，需要更新
		} else {
			r = mid - 1
		}
	}
	return l
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 { return []int {-1, -1} }
	left := findFirstPos(nums, target)
	if left == -1 {
		return []int{-1, -1}
	}
	right := findLastPos(nums, target)
	return []int{left, right}
}