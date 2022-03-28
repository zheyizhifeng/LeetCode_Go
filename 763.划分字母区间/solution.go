// @algorithm @lc id=768 lang=golang 
// @title partition-labels


package main
// @test("ababcbacadefegdehijhklij")=[9,7,8]
// @test("eccbbbbdec")=[10]
func partitionLabels(s string) []int {
	lastIndexMap := make(map[rune]int)
	res := []int{}
	maxIndex := -1
	prevLength := 0
	for i, c := range s {
		lastIndexMap[c] = i
	}
	// lastIndexMap := make(map[byte]int)
	// length := len(s)
	// for i:=0;i<length;i++ {
	// 	lastIndexMap[s[i]] = i
	// }
	for i, c := range s {
		lastIndex := lastIndexMap[c]
		// * math.Max输入为float64类型
		// * maxIndex := math.Max(maxIndex, lastIndex)
		if lastIndex > maxIndex {
			maxIndex = lastIndex
		}
		if i == maxIndex {
			partLength := maxIndex + 1 - prevLength
			res = append(res, partLength)
			prevLength += partLength
		}
	}
	return res
}