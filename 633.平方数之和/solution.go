// @algorithm @lc id=633 lang=golang 
// @title sum-of-square-numbers

package main
// import "math"
// @test(5)=true
// @test(3)=false
func judgeSquareSum(c int) bool {
	left, right := 0, int(math.Sqrt(float64(c)))
	for left <= right {
		sum := left * left + right * right
		if (sum == c) {
			return true
		} else if sum < c {
			left++
		} else {
			right--
		}
	}
	return false
}