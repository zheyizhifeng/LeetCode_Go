// @algorithm @lc id=88 lang=golang
// @title merge-sorted-array

package main

// @test([1,2,3,0,0,0],3,[2,5,6],3)=[1,2,2,3,5,6]
// @test([1],1,[],0)=[1]
// @test([0],0,[1],1)=[1]
func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1
	// * i + j + 1即为插入位置
	for j >= 0 {
		if i >= 0 && nums1[i] >= nums2[j] {
			nums1[i+j+1] = nums1[i]
			i--
		} else {
			nums1[i+j+1] = nums2[j]
			j--
		}
	}
}
