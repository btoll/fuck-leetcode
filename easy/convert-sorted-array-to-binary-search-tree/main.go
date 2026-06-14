package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := len(nums) / 2
	return &TreeNode{
		Val:   nums[root],
		Left:  sortedArrayToBST(nums[:root]),
		Right: sortedArrayToBST(nums[root+1:]),
	}
}

// Can be:
//
//		  0
//	     / \
//	   -3   9
//	   /   /
//	 -10  5
//
// Or:
//
//		  0
//	     / \
//	   -10  5
//	      \  \
//	      -3  9

func main() {
	tests := [][]int{
		{-10, -3, 0, 5, 9},
		{1, 3},
	}
	for _, test := range tests {
		fmt.Println(sortedArrayToBST(test))
	}
}
