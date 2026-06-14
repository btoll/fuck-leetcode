package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Given the root of a binary tree and an integer targetSum, return true if the tree has
// a root-to-leaf path such that adding up all the values along the path equals targetSum.

type TreeNode[T constraints.Integer] struct {
	Val   T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

type Queue[T comparable] []T

func (q *Queue[T]) Empty() bool {
	return len(*q) == 0
}

func (q *Queue[T]) Enqueue(v T) {
	*q = append(*q, v)
}

func (q *Queue[T]) Dequeue() T {
	v := (*q)[0]
	*q = (*q)[1:]
	return v
}

func makeBinaryTreeBFS[T constraints.Integer](nodes []T, sentinel T) *TreeNode[T] {
	root := &TreeNode[T]{Val: nodes[0]}
	q := &Queue[*TreeNode[T]]{root}
	indices := &Queue[int]{0}
	for !q.Empty() {
		node := q.Dequeue()
		index := indices.Dequeue()

		leftNodeIndex := 2*index + 1
		rightNodeIndex := 2*index + 2

		if leftNodeIndex < len(nodes) && nodes[leftNodeIndex] != sentinel {
			node.Left = &TreeNode[T]{Val: nodes[leftNodeIndex]}
			q.Enqueue(node.Left)
			indices.Enqueue(leftNodeIndex)
		}

		if rightNodeIndex < len(nodes) && nodes[rightNodeIndex] != sentinel {
			node.Right = &TreeNode[T]{Val: nodes[rightNodeIndex]}
			q.Enqueue(node.Right)
			indices.Enqueue(rightNodeIndex)
		}
	}
	return root
}

func hasPathSum[T constraints.Integer](root *TreeNode[T], targetSum T) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val
	if root.Left == nil && root.Right == nil { // leaf node
		return targetSum == 0
	}
	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}

func main() {
	nodes := []int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, -1, 1}
	fmt.Println(
		hasPathSum(
			makeBinaryTreeBFS(nodes, -1),
			22,
		))
}
