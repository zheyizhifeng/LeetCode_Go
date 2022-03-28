// @algorithm @lc id=1185 lang=golang 
// @title find-in-mountain-array


package main
// @test([1,2,3,4,5,3,1],3)=2
// @test([0,1,2,4,2,1],3)=-1
/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

func findPeakIndex(mountainArr *MountainArray) int {
	left, right := 0, mountainArr.length() - 1
	for left < right {
		mid := left + (right - left + 1) / 2
		if mountainArr.get(mid) > mountainArr.get(mid - 1) {
			// * 上升
			left = mid // * 需要变动mid取值
		} else {
			// * 下降
			right = mid - 1
		}
	}
	return left
}
func binarySearch(mountainArr *MountainArray, left, right, target int, isAscend bool) int {
	for left < right {
		mid := left + ((right - left) / 2)
		midV := mountainArr.get(mid)
		if midV == target { return mid }
		if midV < target {
			if isAscend {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if isAscend {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	if mountainArr.get(left) == target { return left }
	return -1
}
func findInMountainArray(target int, mountainArr *MountainArray) int {
    length := mountainArr.length()
		peakIndex := findPeakIndex(mountainArr)
		if mountainArr.get(peakIndex) == target { return peakIndex }
		ans := -1
		if target >= mountainArr.get(0) && target <= mountainArr.get(peakIndex - 1) {
			ans = binarySearch(mountainArr, 0, peakIndex - 1, target, true)
			if ans != -1 { return ans }
		}
		if target >= mountainArr.get(length - 1) && target <= mountainArr.get(peakIndex + 1) {
			ans = binarySearch(mountainArr, peakIndex + 1, length - 1, target, false)
			if ans != -1 { return ans }
		}
		return ans
}