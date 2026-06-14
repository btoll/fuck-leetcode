package main

import (
	"fmt"
	"strings"
)

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	bRow := make([]bool, len(matrix))
	bCol := make([]bool, len(matrix[0]))

	for i := range matrix {
		for j := range matrix[0] {
			if matrix[i][j] == 0 {
				bRow[i] = true
				bCol[j] = true
			}
		}
	}

	for i, br := range bRow {
		if br {
			for k := range matrix[i] {
				matrix[i][k] = 0
			}
		}
	}

	for i, bc := range bCol {
		if bc {
			for k := range matrix {
				matrix[k][i] = 0
			}
		}
	}
}

func prettyPrint(matrix [][]int) {
	builder := strings.Builder{}
	for _, row := range matrix {
		for j := range row {
			fmt.Fprintf(&builder, "%d", row[j])
		}
		builder.WriteByte('\n')
	}
	fmt.Println(builder.String())
}

func main() {
	tests := [][][]int{
		{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
		{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}},
	}
	for _, test := range tests {
		setZeroes(test)
	}
}
