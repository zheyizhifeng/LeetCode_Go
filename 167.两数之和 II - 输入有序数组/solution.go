// @algorithm @lc id=167 lang=golang 
// @title two-sum-ii-input-array-is-sorted


package main
// @test([2,7,11,15],9)=[1,2]
// @test([2,3,4],6)=[1,3]
// @test([-1,0],-1)=[1,2]
func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers) - 1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum > target {
			right--
		} else {
			left++
		}
	}
	return []int{}
}