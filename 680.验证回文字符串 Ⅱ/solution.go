// @algorithm @lc id=680 lang=golang 
// @title valid-palindrome-ii


package main
// @test("aba")=true
// @test("abca")=true
// @test("abc")=false
func isPalindrome(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}
func validPalindrome(s string) bool {
	left, right := 0, len(s) - 1
	for left < right {
		if (s[left] == s[right]) {
			left++
			right--
		} else {
			return isPalindrome(s, left + 1, right) || isPalindrome(s, left, right - 1)
		}
	}
	return true
}