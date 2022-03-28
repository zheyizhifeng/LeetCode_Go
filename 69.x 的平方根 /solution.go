// @algorithm @lc id=69 lang=golang 
// @title sqrtx


package main
// @test(4)=2
// @test(8)=2
func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	left, right := 0, x;
	ans := 0
	for left <= right {
		mid := (left + right) / 2
		if mid * mid <= x {
			ans = mid;
			left = mid + 1;
		} else {
			right = mid - 1;
		}
	}
	return ans
}