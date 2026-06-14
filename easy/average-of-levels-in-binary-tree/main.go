package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func averageOfLevels(root *TreeNode) []float64 {
//	if root == nil {
//		return nil
//	}
//	var dfs func(*TreeNode, int)
//	var levels [][]int
//	dfs = func(root *TreeNode, level int) {
//		if root == nil {
//			return
//		}
//		if len(levels) <= level {
//			levels = append(levels, []int{})
//		}
//		levels[level] = append(levels[level], root.Val)
//		dfs(root.Left, level+1)
//		dfs(root.Right, level+1)
//	}
//	dfs(root, 0)
//	averages := make([]float64, 0, len(levels))
//	for _, level := range levels {
//		var acc float64 = 0
//		for _, v := range level {
//			acc += float64(v)
//		}
//		averages = append(averages, acc/float64(len(level)))
//	}
//	return averages
//}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	nodes := []*TreeNode{root}
	level := 0
	res := []float64{}
	for len(nodes) > 0 {
		l := len(nodes)
		i := 0
		var acc float64 = 0
		for i < l {
			node := nodes[i]
			acc += float64(node.Val)
			if node.Left != nil {
				nodes = append(nodes, node.Left)
			}
			if node.Right != nil {
				nodes = append(nodes, node.Right)
			}
			i++
		}
		nodes = nodes[i:]
		res = append(res, acc/float64(l))
		level++
	}
	return res
}

func makeTree(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := preorder[0]
	inorderIndex := 0
	for i, v := range inorder {
		if root == v {
			inorderIndex = i
			break
		}
	}
	return &TreeNode{
		Val:   root,
		Left:  makeTree(preorder[1:inorderIndex+1], inorder[:inorderIndex]),
		Right: makeTree(preorder[inorderIndex+1:], inorder[inorderIndex+1:]),
	}
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var nodes []int
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		nodes = append(nodes, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return nodes
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := makeTree(preorder, inorder)
	fmt.Printf("root=%#v\n", root)
	fmt.Println(preorderTraversal(root))
	fmt.Println(averageOfLevels(root))
}
