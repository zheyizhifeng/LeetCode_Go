// @algorithm @lc id=417 lang=golang
// @title pacific-atlantic-water-flow

package main

// @test([[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]])=[[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]]
// @test([[2,1],[1,2]])=[[0,0],[0,1],[1,0],[1,1]]
func dfs(heights [][]int, canReach [][]bool, i, j int) {
	m, n := len(heights), len(heights[0])
	if canReach[i][j] {
		return
	}
	canReach[i][j] = true
	direction := [4][2]int{
		{i - 1, j},
		{i + 1, j},
		{i, j - 1},
		{i, j + 1},
	}
	for _, pos := range direction {
		newI, newJ := pos[0], pos[1]
		if newI >= 0 && newI < m && newJ >= 0 && newJ < n && heights[newI][newJ] >= heights[i][j] {
			dfs(heights, canReach, newI, newJ)
		}
	}
}
func pacificAtlantic(heights [][]int) [][]int {
	ans := [][]int{}
	m := len(heights)
	if m == 0 {
		return ans
	}
	n := len(heights[0])
	if n == 0 {
		return ans
	}
	pacific, atlantic := make([][]bool, m), make([][]bool, m)

	for i := 0; i < m; i++ {
		pacific[i] = make([]bool, n)
		atlantic[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		dfs(heights, pacific, i, 0)
		dfs(heights, atlantic, i, n-1)
	}

	for j := 0; j < n; j++ {
		dfs(heights, pacific, 0, j)
		dfs(heights, atlantic, m-1, j)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if pacific[i][j] && atlantic[i][j] {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return ans
}
