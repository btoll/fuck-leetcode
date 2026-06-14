package main

import "fmt"

// Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree
// and inorder is the inorder traversal of the same tree, construct and return the binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// preorder=[1 2 4 8 5 3 6 7]
// inorder=[8 4 2 5 1 6 3 7]
// preorder=[2 4 8 5]
// inorder=[8 4 2 5]
// preorder=[4 8]
// inorder=[8 4]
// preorder=[8]
// inorder=[8]
// preorder=[5]
// inorder=[5]
// preorder=[3 6 7]
// inorder=[6 3 7]
// preorder=[6]
// inorder=[6]
// preorder=[7]
// inorder=[7]
// &{1 0x37488c8f21c8 0x37488c8f22a0}
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 || len(preorder) == 0 {
		return nil
	}

	fmt.Printf("preorder=%+v\n", preorder)
	fmt.Printf("inorder=%+v\n", inorder)
	root := preorder[0]
	var inorderIndex int
	for i, val := range inorder {
		if root == val {
			inorderIndex = i
			break
		}
	}

	return &TreeNode{
		Val:   root,
		Left:  buildTree(preorder[1:inorderIndex+1], inorder[:inorderIndex]),
		Right: buildTree(preorder[inorderIndex+1:], inorder[inorderIndex+1:]),
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
	// preorder := []int{3, 9, 20, 15, 7}
	// inorder := []int{9, 3, 15, 20, 7}
	preorder := []int{1, 2, 4, 8, 5, 3, 6, 7}
	inorder := []int{8, 4, 2, 5, 1, 6, 3, 7}
	fmt.Println(buildTree(preorder, inorder))
}
