// @algorithm @lc id=7 lang=golang
// @title reverse-integer

package main

import (
	"math"
)

// @test(123)=321
// @test(-123)=-321
// @test(120)=21
func reverse(x int) (rev int) {
	for x != 0 {
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		rev = rev*10 + digit
	}
	return
}
