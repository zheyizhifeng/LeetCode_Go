// @algorithm @lc id=3 lang=golang
// @title longest-substring-without-repeating-characters

package main

// @test("abcabcbb")=3
// @test("bbbbb")=1
// @test("pwwkew")=3
func lengthOfLongestSubstring(s string) int {
	l, r := 0, 0
	max := 0
	window := make(map[byte]int)
	for i := range s {
		if _, ok := window[s[i]]; ok {
			last := window[s[i]]
			for j := l; j <= last; j++ {
				delete(window, s[j])
			}
			l = last + 1
		}
		r++
		window[s[i]] = i
		if r-l > max {
			max = r - l
		}
	}
	return max
}
