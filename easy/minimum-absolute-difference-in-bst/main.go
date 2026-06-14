package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func getMinimumDifference(root *TreeNode) int {
//	inorder := inorderTraversal(root)
//	if len(inorder) < 2 {
//		return 0
//	}
//	minDiff := math.MaxInt
//	for i := 1; i < len(inorder); i++ {
//		minDiff = min(minDiff, inorder[i]-inorder[i-1])
//	}
//	return minDiff
//}

//func getMinimumDifference(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	minDiff := math.MaxInt
//	minDiff = min(minDiff, root.Val-getMinimumDifference(root.Left))
//	minDiff = min(minDiff, root.Val-getMinimumDifference(root.Right))
//	return minDiff
//}

func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}
	minDiff := math.MaxInt
	var prev *int
	var dfs func(*TreeNode)
	dfs = func(n *TreeNode) {
		if n == nil {
			return
		}
		dfs(n.Left)
		if prev != nil {
			diff := n.Val - *prev
			if diff < minDiff {
				minDiff = diff
			}
		}
		prev = &n.Val
		dfs(n.Right)
	}
	dfs(root)
	return minDiff
}

func inorderTraversal(node *TreeNode) []int {
	if node == nil {
		return nil
	}
	var dfs func(*TreeNode)
	nodes := []int{}
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		nodes = append(nodes, root.Val)
		dfs(root.Right)
	}
	dfs(node)
	return nodes
}

//func makeTree(preorder, inorder []int) *TreeNode {
//	if len(preorder) == 0 || len(inorder) == 0 {
//		return nil
//	}
//	root := preorder[0]
//	inorderIndex := 0
//	for i, v := range inorder {
//		if v == root {
//			inorderIndex = i
//			break
//		}
//	}
//	return &TreeNode{
//		Val:   root,
//		Left:  makeTree(preorder[1:inorderIndex+1], inorder[:inorderIndex]),
//		Right: makeTree(preorder[inorderIndex+1:], inorder[inorderIndex+1:]),
//	}
//}

//func preorderTraversal(root *TreeNode) []int {
//	if root == nil {
//		return nil
//	}
//	var dfs func(*TreeNode)
//	nodes := []int{}
//	dfs = func(root *TreeNode) {
//		if root == nil {
//			return
//		}
//		nodes = append(nodes, root.Val)
//		dfs(root.Left)
//		dfs(root.Right)
//	}
//	dfs(root)
//	return nodes
//}

func main() {
	//	tests := [][]int{
	//		{4, 2, 6, 1, 3},
	//		{1, 0, 48, -1, -1, 12, 49},
	//	}
	tests := []*TreeNode{
		{
			Val: 4,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
		{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 48,
				Left: &TreeNode{
					Val: 12,
				},
				Right: &TreeNode{
					Val: 49,
				},
			},
		},
	}
	for _, test := range tests {
		//		preorder := preorderTraversal(test)
		//		inorder := inorderTraversal(test)
		//		fmt.Printf("\npreorder=%#v\n", preorder)
		//		fmt.Printf("preorder=%#v\n", preorderTraversal(makeTree(preorder, inorder)))
		//		fmt.Printf("inorder=%#v\n", inorder)
		//		fmt.Printf("inorder=%#v\n", inorderTraversal(makeTree(preorder, inorder)))
		fmt.Println(getMinimumDifference(test))
	}
}
