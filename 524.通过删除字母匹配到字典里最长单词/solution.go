// @algorithm @lc id=524 lang=golang
// @title longest-word-in-dictionary-through-deleting

package main

// @test("abpcplea",["ale","apple","monkey","plea"])="apple"
// @test("abpcplea",["a","b","c"])="a"
func findLongestWord(s string, dictionary []string) string {
  res := ""
  len1 := len(s)
  for _, word := range dictionary {
    len2 := len(word)
    i, j := 0, 0
    for i < len1 && j < len2 {
      if s[i] == word[j] {
        j++
      }
      i++
    }
    if j == len2 {
      len3 := len(res)
      if len2 > len3 {
        res = word
      } else if len2 == len3 && word < res {
        res = word
      }
    }
  }
  return res
}
