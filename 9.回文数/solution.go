// @algorithm @lc id=9 lang=golang 
// @title palindrome-number


package main
// @test(121)=true
// @test(-121)=false
// @test(10)=false
// @test(-101)=false
func isPalindrome(x int) bool {
	if x < 0 || (x % 10 == 0 && x != 0) {
		return false;
	} else {
		right := 0;
		for x > right {
			right = right * 10 + x % 10;
			x = x / 10;
		}
		return x == right || x == right / 10;
	}
}