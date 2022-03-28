// @algorithm @lc id=547 lang=golang
// @title number-of-provinces

package main

// @test([[1,1,0],[1,1,0],[0,0,1]])=2
// @test([[1,0,0],[0,1,0],[0,0,1]])=3
func dfs(isConnected [][]int, i int, visited []bool) {
	for j := 0; j < len(isConnected); j++ {
		if isConnected[i][j] == 1 && !visited[j] {
			visited[j] = true
			dfs(isConnected, j, visited)
		}
	}
}


func findCircleNum(isConnected [][]int) int {
	cityLen := len(isConnected)
	visited := make([]bool, cityLen)
	provinces := 0
	for i := 0; i < cityLen; i++ {
		if !visited[i] {
			dfs(isConnected, i, visited)
			provinces++
		}
	}
	return provinces
}
