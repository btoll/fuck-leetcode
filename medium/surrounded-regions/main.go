package main

import "fmt"

// The idea is to do a DFS from all of the border rows and columns where the cell == 'O'.
// Mark everything is can reach as 'S', for "safe".
// Iterate over the board again, first changing all existing 'O's to an 'X', and then
// lastly flipping all 'S's back to an 'O'.  Order matters here.
func solve(board [][]byte) {
	// Check the first and last columns.
	for i := range board {
		if board[i][0] == 'O' {
			dfs(board, i, 0)
		}
		if board[i][len(board[0])-1] == 'O' {
			dfs(board, i, len(board[0])-1)
		}
	}
	// Check the first and last rows.
	for j := range board[0] {
		if board[0][j] == 'O' {
			dfs(board, 0, j)
		}
		if board[len(board)-1][j] == 'O' {
			dfs(board, len(board)-1, j)
		}
	}
	// Flip back to 'O'.
	for i := range board {
		for j := range board[0] {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == 'S' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfs(board [][]byte, i, j int) {
	if i < 0 ||
		j < 0 ||
		i >= len(board) ||
		j >= len(board[0]) ||
		board[i][j] != 'O' {
		return
	}
	board[i][j] = 'S'
	dfs(board, i-1, j)
	dfs(board, i, j+1)
	dfs(board, i+1, j)
	dfs(board, i, j-1)
}

func main() {
	tests := [][][]byte{
		{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}},
		{{'X', 'X', 'X'}, {'X', 'O', 'X'}, {'X', 'X', 'X'}},
		{{'O', 'O', 'O'}, {'O', 'O', 'O'}, {'O', 'O', 'O'}},
	}
	for _, test := range tests {
		solve(test)
		fmt.Printf("board=%+v\n", test)
	}
}
