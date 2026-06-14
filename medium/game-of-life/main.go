package main

import "fmt"

func gameOfLife(board [][]int) {

}

func main() {
	board := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}

	fmt.Println(gameOfLife(board))
}
