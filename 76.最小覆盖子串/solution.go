// @algorithm @lc id=76 lang=golang
// @title minimum-window-substring

package main

// @test("ADOBECODEBANC","ABC")="BANC"
// @test("a","a")="a"
// @test("a","aa")=""
// @test("a","b")=""
func minWindow(s string, t string) string {
	lenS := len(s)
	if len(t) > lenS {
		return ""
	}
	window, need := make(map[string]int), make(map[string]int)
	start, minLength := 0, lenS + 1
	l, r := 0, 0
	valid := 0
	for _, v := range t {
		v := string(v)
		need[v]++
	}
	for r <= lenS {
		if valid == len(need) {
			leftChar := string(s[l])
			if r - l < minLength {
				minLength = r - l
				start = l
			}
			if _, ok := window[leftChar]; ok {
				if window[leftChar] == need[leftChar] {
					valid--
				}
				window[leftChar] -= 1
			}
			l++
		} else {
			if r < lenS {
				rightChar := string(s[r])
				if _, ok := need[rightChar]; ok {
					window[rightChar]++
					if window[rightChar] == need[rightChar] {
						valid++
					}
				}
			}
			r++
		}
	}
	if (minLength == lenS + 1) {
		return ""
	}
	return s[start: start + minLength]
}