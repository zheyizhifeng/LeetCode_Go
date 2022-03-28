// @algorithm @lc id=28 lang=golang 
// @title implement-strstr

package main
// @test("hello","ll")=2
// @test("aaaaa","bba")=-1
// @test("","")=0
func getPMT(pattern string) []int {
	i, j := 1, 0;
	length := len(pattern);
	PMT := make([]int, length);
	for i < length {
		if pattern[i] == pattern[j] {
			j++;
			PMT[i] = j;
			i++
		} else {
			if (j > 0) {
				j = PMT[j-1];
			} else {
				i++;
			}
		}
	}
	return PMT;
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 || needle == haystack {
		return 0
	}
	PMT := getPMT(needle);
	i, j := 0, 0;
	for i < len(haystack) {
		if haystack[i] == needle[j] {
			i++;
			j++;
			if (j == len(needle)) {
				return i - j // or i - len(needle)
			}
		} else {
			if j > 0 {
				j = PMT[j-1];
			} else {
				i++;
			}
		}
	}
	return -1;
}