package main

import "fmt"

// Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.
// An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all
// four edges of the grid are all surrounded by water.
func numIslands(grid [][]int) int {
	var visited int
	for rowIndex, row := range grid {
		for columnIndex, column := range row {
			if column == 1 {
				visited += 1
				dfs(grid, rowIndex, columnIndex)
			}
		}
	}
	return visited
}

func dfs(grid [][]int, rowIndex, columnIndex int) {
	if rowIndex < 0 ||
		columnIndex < 0 ||
		rowIndex > len(grid)-1 ||
		columnIndex > len(grid[0])-1 ||
		grid[rowIndex][columnIndex] != 1 {
		return
	}

	grid[rowIndex][columnIndex] = 0
	top := []int{rowIndex - 1, columnIndex}
	right := []int{rowIndex, columnIndex + 1}
	bottom := []int{rowIndex + 1, columnIndex}
	left := []int{rowIndex, columnIndex - 1}

	dfs(grid, top[0], top[1])
	dfs(grid, right[0], right[1])
	dfs(grid, bottom[0], bottom[1])
	dfs(grid, left[0], left[1])
}

func main() {
	//	grid := [][]byte{
	//		{'1', '1', '1', '1', '0'},
	//		{'1', '1', '0', '1', '0'},
	//		{'1', '1', '0', '0', '0'},
	//		{'0', '0', '0', '0', '0'},
	//	}
	grid := [][]int{
		{1, 1, 1, 1, 0},
		{0, 0, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 1},
	}
	//	grid := [][]int{
	//		{1, 1, 0, 1},
	//		{1, 0, 0, 0},
	//		{0, 1, 1, 0},
	//		{1, 0, 1, 1},
	//	}

	fmt.Println(numIslands(grid))
}
