// @algorithm @lc id=2106 lang=golang 
// @title find-greatest-common-divisor-of-array


package main
// @test([2,5,6,9,10])=2
// @test([7,5,6,8,3])=1
// @test([3,3])=3
// @test([10,6,9])=2
func findGCD(nums []int) int {
	minimum, maximum := 1001, 0
	for _, n := range nums {
		if n < minimum {
			minimum = n
		}
		if n > maximum {
			maximum = n
		}
	}
	for minimum != 0 {
		maximum, minimum = minimum, maximum % minimum
	}
	return maximum
}