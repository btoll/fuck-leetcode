package main

import "fmt"

// Given two integer arrays inorder and postorder where inorder is the inorder traversal of a binary tree
// and postorder is the postorder traversal of the same tree, construct and return the binary tree.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	root := postorder[len(postorder)-1]
	var inorderIndex int
	for i, val := range inorder {
		if root == val {
			inorderIndex = i
			break
		}
	}

	return &TreeNode{
		Val:   root,
		Left:  buildTree(inorder[:inorderIndex], postorder[:inorderIndex]),
		Right: buildTree(inorder[inorderIndex+1:], postorder[inorderIndex:len(postorder)-1]),
	}
}

// preorder  - root -> left -> right
// inorder   - left -> root -> right
// postorder - left -> right -> root

//           1
//         /   \
//        2     3
//       / \   / \
//      4   5 6   7
//     /
//    8

//       3
//      / \
//     9   20
//        /  \
//       15   7

func main() {
	//	preorder := []int{3, 9, 20, 15, 7}
	//	inorder := []int{9, 3, 15, 20, 7}
	//	postorder := []int{9, 15, 7, 20, 3}
	inorder := []int{8, 4, 2, 5, 1, 6, 3, 7}
	postorder := []int{8, 4, 5, 2, 6, 7, 3, 1}
	fmt.Println(buildTree(inorder, postorder))
}
