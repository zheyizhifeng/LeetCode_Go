// @algorithm @lc id=10 lang=golang
// @title regular-expression-matching

package main

// @test("aa","a")=false
// @test("aa","a*")=true
// @test("ab",".*")=true
// @test("aab", "c*a*b")=true
func match(a, b byte) bool {
	return a == b || b == '.'
}

func isMatch(s string, p string) bool {
	lenS, lenP := len(s), len(p)
	dp := make([][]bool, lenS+1)
	for i := range dp {
		dp[i] = make([]bool, lenP+1)
	}
	dp[0][0] = true
	// badcase
	for j := 1; j <= lenP; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2];
		}
	}
	for i := 1; i <= lenS; i++ {
		for j := 1; j <= lenP; j++ {
			if match((s[i-1]), (p[j-1])) {
				dp[i][j] = dp[i-1][j-1]
			} else if (p[j-1]) == '*' {
				if match((s[i-1]), (p[j-2])) {
					dp[i][j] = dp[i][j-2] || dp[i-1][j-2] || dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-2]
				}
			}
		}
	}
	return dp[lenS][lenP]
}
