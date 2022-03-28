// @algorithm @lc id=406 lang=golang
// @title queue-reconstruction-by-height

package main

// import "sort"

// @test([[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]])=[[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]
// @test([[6,0],[5,0],[4,0],[3,2],[2,2],[1,4]])=[[4,0],[5,0],[2,2],[3,2],[1,4],[6,0]]

/*
func insert(arr [][]int, index int, element []int) [][]int {
	arr = append(arr, []int{})
	for i := len(arr) - 1; i > index; i-- {
		arr[i] = arr[i-1]
	}
	arr[index] = element
	return arr
} 
*/
func reconstructQueue(people [][]int) [][]int {
	res := make([][]int, len(people))
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		if a[0] == b[0] {
			return a[1] < b[1]
		} else {
			return a[0] > b[0]
		}
	})
	for _, person := range people {
		// res = insert(res, person[1], person)
		copy(res[person[1] + 1:], res[person[1]:])
		res[person[1]] = person
	}
	return res
}
