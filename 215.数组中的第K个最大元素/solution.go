// @algorithm @lc id=215 lang=golang
// @title kth-largest-element-in-an-array

package main

import "math/rand"

// @test([3,2,1,5,6,4],2)=5
// @test([3,2,3,1,2,4,5,5,6],4)=4
func random(l, r int) int {
	return rand.Int() % (r-l) + l
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func partition(arr []int, l, r int) int {
	pivotIndex := random(l ,r)
	pivot := arr[pivotIndex]
	swap(arr, l, pivotIndex)
	less_than_pivot_index := l
	for i := l + 1; i<= r; i++ {
		if arr[i] < pivot {
			less_than_pivot_index++
			swap(arr, i, less_than_pivot_index)
		}
	}
	swap(arr, l, less_than_pivot_index)
	return less_than_pivot_index
}

func findKthLargest(nums []int, k int) int {
	length := len(nums)
	l, r := 0, length - 1
	target := length - k
	for l < r {
		partitionIndex := partition(nums, l, r)
		if partitionIndex == target {
			return nums[partitionIndex]
		} else if (partitionIndex < target) {
			l = partitionIndex + 1
		} else {
			r = partitionIndex - 1
		}
	}
	return nums[l]
}
