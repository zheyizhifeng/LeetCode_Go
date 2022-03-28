// @algorithm @lc id=4 lang=golang
// @title median-of-two-sorted-arrays

package main

// @test([1,3],[2])=2.00000
// @test([1,2],[3,4])=2.50000
func findKthElement(nums1 []int, nums2 []int, k int) int {
	i, j, len1, len2 := 0, 0, len(nums1), len(nums2)
	for {
		if i == len1 {
			return nums2[j+k-1]
		} else if j == len2 {
			return nums1[i+k-1]
		} else if k == 1 {
			if nums1[i] < nums2[j] {
				return nums1[i]
			}
			return nums2[j]
		}
		half := k / 2
		newI, newJ := len1-1, len2-1
		if i+half-1 < len1-1 {
			newI = i + half - 1
		}
		if j+half-1 < len2-1 {
			newJ = j + half - 1
		}
		if nums1[newI] >= nums2[newJ] {
			k = k - (newJ - j + 1)
			j = newJ + 1
		} else {
			k = k - (newI - i + 1)
			i = newI + 1
		}
	}
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	medianLeft, medianRight := totalLength/2, (totalLength-1)/2
	if medianLeft == medianRight {
		return float64(findKthElement(nums1, nums2, medianLeft+1))
	} else {
		return float64(findKthElement(nums1, nums2, medianLeft+1)+findKthElement(nums1, nums2, medianRight+1)) / 2.0
	}
}
