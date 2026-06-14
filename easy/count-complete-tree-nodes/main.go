package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

func makeCompleteBinaryTree(nodes []int) *TreeNode {
	if len(nodes) == 0 {
		return nil
	}
	treeNodes := make([]*TreeNode, len(nodes))
	for i, v := range nodes {
		treeNodes[i] = &TreeNode{Val: v}
	}
	for i := range nodes {
		left := 2*i + 1
		right := 2*i + 2
		if left < len(nodes) {
			treeNodes[i].Left = treeNodes[left]
		}
		if right < len(nodes) {
			treeNodes[i].Right = treeNodes[right]
		}
	}
	return treeNodes[0]
}

//func makeCompleteBinaryTree(nodes []int) *TreeNode {
//	if len(nodes) == 0 {
//		return nil
//	}
//	root := &TreeNode{Val: nodes[0]}
//	q := []*TreeNode{root}
//	i := 1
//	for len(q) > 0 && i < len(nodes) {
//		node := q[0]
//		q = q[1:]
//		node.Left = &TreeNode{Val: nodes[i]}
//		q = append(q, node.Left)
//		i++
//		if i < len(nodes) {
//			node.Right = &TreeNode{Val: nodes[i]}
//			q = append(q, node.Right)
//			i++
//		}
//	}
//	return root
//}

func bfs(root *TreeNode) []int {
	q := []*TreeNode{root}
	nodes := []int{root.Val}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node.Left != nil {
			nodes = append(nodes, node.Left.Val)
			q = append(q, node.Left)
		}
		if node.Right != nil {
			nodes = append(nodes, node.Right.Val)
			q = append(q, node.Right)
		}
	}
	return nodes
}

func inorder(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.Val)
	inorder(node.Left)
	inorder(node.Right)
}

func main() {
	nodes := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("nodes=%#v\n", nodes)
	root := makeCompleteBinaryTree(nodes)
	fmt.Println(root)
	inorder(root)
	fmt.Printf("bfs=%#v\n", bfs(root))
	fmt.Println(countNodes(root))
}
