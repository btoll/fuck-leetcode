package main

import (
	"fmt"
	"strings"
)

func prettyPrint(matrix [][]int, print bool) {
	if !print {
		return
	}
	builder := strings.Builder{}
	for i, row := range matrix {
		for j := range row {
			fmt.Fprintf(&builder, "%d ", matrix[i][j])
		}
		builder.WriteByte('\n')
	}
	fmt.Println(builder.String())
}

func rotate(matrix [][]int) {
	prettyPrint(matrix, true)
	// Transpose.
	l := len(matrix)
	for i := range l {
		for j := i + 1; j < l; j++ {
			//			fmt.Printf("i=%+v\n", i)
			//			fmt.Printf("j=%+v\n", j)
			//			fmt.Printf("matrix[i][j]=%+v\n", matrix[i][j])
			//			fmt.Printf("matrix[j][i]=%+v\n", matrix[j][i])
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
			prettyPrint(matrix, true)
		}
	}

	// Reverse.
	for i := range l {
		for j := 0; j < l/2; j++ {
			matrix[i][l-j-1], matrix[i][j] = matrix[i][j], matrix[i][l-j-1]
		}
	}
	prettyPrint(matrix, true)
}

//func rotate(matrix [][]int) {
//	n := len(matrix)
//	if n == 0 || len(matrix[0]) != n {
//		return // requires N×N
//	}
//
//	for layer := 0; layer < n/2; layer++ {
//		first := layer
//		last := n - 1 - layer
//		for i := first; i < last; i++ {
//			offset := i - first
//			// save top
//			top := matrix[first][i]
//			// left -> top
//			matrix[first][i] = matrix[last-offset][first]
//			// bottom -> left
//			matrix[last-offset][first] = matrix[last][last-offset]
//			// right -> bottom
//			matrix[last][last-offset] = matrix[i][last]
//			// top -> right
//			matrix[i][last] = top
//		}
//	}
//}

// Original:
// 1 2 3
// 4 5 6
// 7 8 9
//
// After transpose:
// 1 4 7
// 2 5 8
// 3 6 9
//
// After reversing each row (90° clockwise):
// 7 4 1
// 8 5 2
// 9 6 3
//
// After reversing each row (90° counter-clockwise):
// 3 6 9
// 2 5 8
// 1 4 7

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// {1, 2, 3, 4}
	// {4, 3, 2, 1}

	// {1, 2, 3, 4, 5, 6}
	// {6, 5, 4, 3, 2, 1}

	fmt.Printf("matrix=%+v\n", matrix)
	rotate(matrix)
	fmt.Printf("matrix=%+v\n", matrix)
}
