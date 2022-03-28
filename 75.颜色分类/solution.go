// @algorithm @lc id=75 lang=golang 
// @title sort-colors


package main
// @test([2,0,2,1,1,0])=[0,0,1,1,2,2]
// @test([2,0,1])=[0,1,2]
func sortColors(nums []int)  {
	zero, one, two := 0, 0, 0
	for _, n := range nums {
		if n == 0 {
			zero++
		} else if n == 1 {
			one++
		} else {
			two++
		}
	}
	i := 0
	for i < len(nums) {
		for j := 0; j < zero; j++ {
			nums[i] = 0
			i++
		}
		for j := 0; j < one; j++ {
			nums[i] = 1
			i++
		}
		for j := 0; j < two; j++ {
			nums[i] = 2
			i++
		}
	}
}