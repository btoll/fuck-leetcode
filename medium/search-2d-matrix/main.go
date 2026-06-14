package main

import "fmt"

// You are given an m x n integer matrix matrix with the following two properties:
//
//	Each row is sorted in non-decreasing order.
//	The first integer of each row is greater than the last integer of the previous row.
//
// Given an integer target, return true if target is in matrix or false otherwise.
// You must write a solution in O(log(m * n)) time complexity.
func searchMatrix(matrix [][]int, target int) bool {
	return search(matrix, 0, len(matrix)*len(matrix[0])-1, target)
}

func search(matrix [][]int, p, r, target int) bool {
	if p > r {
		return false
	}
	q := (p + r) / 2
	cell := getCell(matrix, q)
	if target == cell {
		return true
	}
	if target > cell {
		p = q + 1
	} else {
		r = q - 1
	}
	return search(matrix, p, r, target)
}

func getCell(matrix [][]int, q int) int {
	cols := len(matrix[0])
	return matrix[q/cols][q%cols]
}

func main() {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	target := 3
	fmt.Println(searchMatrix(matrix, target))
}
