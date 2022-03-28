// @algorithm @lc id=435 lang=golang
// @title non-overlapping-intervals

package main

// @test([[1,2],[2,3],[3,4],[1,3]])=1
// @test([[1,2],[1,2],[1,2]])=2
// @test([[1,2],[2,3]])=0
// @test([[1,100],[11,22],[1,11],[2,12]])=2
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i int, j int) bool { return intervals[i][1] < intervals[j][1] })
	ans, end := 1, intervals[0][1]

	for _, nums := range intervals[1:] {
		if nums[0] >= end {
			ans++
			end = nums[1]
		}
	}
	return len(intervals) - ans
}
