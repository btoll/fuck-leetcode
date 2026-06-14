package main

import (
	"fmt"
	"strings"
)

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	m := len(matrix)
	n := len(matrix[0])
	nums := make([]int, 0, m*n)
	top, bottom := 0, m-1
	left, right := 0, n-1
	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			nums = append(nums, matrix[top][i])
		}
		top++
		for i := top; i <= bottom; i++ {
			nums = append(nums, matrix[i][right])
		}
		right--
		if top <= bottom {
			for i := right; i >= left; i-- {
				nums = append(nums, matrix[bottom][i])
			}
			bottom--
		}
		if left <= right {
			for i := bottom; i >= top; i-- {
				nums = append(nums, matrix[i][left])
			}
			left++
		}
	}
	return nums
}

func prettyPrint(matrix [][]int) {
	builder := strings.Builder{}
	for _, row := range matrix {
		for j := range row {
			fmt.Fprintf(&builder, "%d ", row[j])
		}
		builder.WriteByte('\n')
	}
	fmt.Println(builder.String())
}

func main() {
	tests := [][][]int{
		{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
	}
	for _, test := range tests {
		prettyPrint(test)
		fmt.Printf("%#v\n\n", spiralOrder(test))
	}
}
