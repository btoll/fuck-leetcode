package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func isValidBST(root *TreeNode) bool {
//	var nodes []int
//	var dfs func(*TreeNode) bool
//	dfs = func(node *TreeNode) bool {
//		if node == nil {
//			return true
//		}
//		if !dfs(node.Left) {
//			return false
//		}
//		if len(nodes) > 0 && node.Val <= nodes[len(nodes)-1] {
//			return false
//		}
//		nodes = append(nodes, node.Val)
//		return dfs(node.Right)
//	}
//	return dfs(root)
//}

func isValidBST(root *TreeNode) bool {
	var prev *TreeNode
	isValid := true
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		// `!isValid` avoids unnecessary traversals once a failure is detected.
		if node == nil || !isValid {
			return
		}
		dfs(node.Left)
		if prev != nil && node.Val <= prev.Val {
			isValid = false
			return
		}
		prev = node
		dfs(node.Right)
	}
	dfs(root)
	return isValid
}

func main() {
	tests := []*TreeNode{
		{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		{
			Val: 1,
			Right: &TreeNode{
				Val: 1,
			},
		},
		{
			Val: 5,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		},
	}
	for _, test := range tests {
		fmt.Println(isValidBST(test))
	}
}
