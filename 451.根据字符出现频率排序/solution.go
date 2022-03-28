// @algorithm @lc id=451 lang=golang
// @title sort-characters-by-frequency

package main

// @test("tree")="eert"
// @test("cccaaa")="aaaccc"
// @test("Aabb")="bbAa"
func frequencySort(s string) string {
	countMap := make(map[byte]int)
	maxCount := 0
	for i := range s {
		countMap[s[i]]++
		if countMap[s[i]] > maxCount {
			maxCount = countMap[s[i]]
		}
	}
	bucket := make([][]byte, maxCount+1)
	for key, value := range countMap {
		bucket[value] = append(bucket[value], key)
	}
	ans := make([]byte, 0, len(s))	
	for i := maxCount; i >= 0; i-- {
		for _, c := range bucket[i] {
			for j := 0; j < i; j++ {
				ans = append(ans, c)
			}
		}
	}
	return string(ans)
}
