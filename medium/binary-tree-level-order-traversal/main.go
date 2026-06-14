package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func makeTree(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := preorder[0]
	inorderIndex := 0
	for index, v := range inorder {
		if root == v {
			inorderIndex = index
			break
		}
	}
	return &TreeNode{
		Val:   root,
		Left:  makeTree(preorder[1:inorderIndex+1], inorder[:inorderIndex]),
		Right: makeTree(preorder[inorderIndex+1:], inorder[inorderIndex+1:]),
	}
}

func inorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	inorderTraversal(root.Left)
	fmt.Println(root.Val)
	inorderTraversal(root.Right)
}

func main() {
	preorder := []int{1, 2, 3, 5, 4}
	inorder := []int{2, 5, 1, 3, 4}
	root := makeTree(preorder, inorder)
	inorderTraversal(root)
}
