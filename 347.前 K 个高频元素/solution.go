// @algorithm @lc id=347 lang=golang
// @title top-k-frequent-elements

package main

// @test([1,1,1,2,2,3],2)=[1,2]
// @test([1],1)=[1]
func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int)
	maxCount := 0
	for _, n := range nums {
		countMap[n] += 1
		if countMap[n] > maxCount {
			maxCount = countMap[n]
		}
	}
	buckets := make([][]int, maxCount + 1)
	for key, value := range countMap {
		buckets[value] = append(buckets[value], key)
	}
	res := []int{}
	for i := maxCount; i >= 0; i-- {
		for _, n := range buckets[i] {
			res = append(res, n)
		}
		if len(res) == k {
			return res
		}
	}
	return []int{}
}
