// @algorithm @lc id=66 lang=golang
// @title plus-one

package main

// @test([1,2,3])=[1,2,4]
// @test([4,3,2,1])=[4,3,2,2]
// @test([9])=[1,0]
func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		bit := (digits[i] + carry) % 10
		carry = (digits[i] + carry) / 10
		digits[i] = bit
	}
	if carry > 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
