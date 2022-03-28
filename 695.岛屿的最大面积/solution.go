// @algorithm @lc id=695 lang=golang
// @title max-area-of-island

package main

// @test([[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]])=6
// @test([[0,0,0,0,0,0,0,0]])=0
func dfs(grid [][]int, i, j int) int {
	rowLen := len(grid)
	columnLen := len(grid[0])
	if i < 0 || j < 0 || i >= rowLen || j >= columnLen || grid[i][j] == 0 {
		return 0
	}
	grid[i][j] = 0
	return 1 + dfs(grid, i, j+1) + dfs(grid, i, j-1) + dfs(grid, i-1, j) + dfs(grid, i+1, j)
}
func maxAreaOfIsland(grid [][]int) int {
	max := 0
	rowLen := len(grid)
	columnLen := len(grid[0])
	for i := 0; i < rowLen; i++ {
		for j := 0; j < columnLen; j++ {
			if grid[i][j] == 1 {
				area := dfs(grid, i, j)
				if area > max {
					max = area
				}
			}
		}
	}
	return max
}
