package main

import (
	"errors"
	"fmt"
)

//Given the root of a binary tree, invert the tree, and return its root.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue[T any] []T

func (q *Queue[T]) Enqueue(v T) {
	*q = append(*q, v)
}

func (q *Queue[T]) Dequeue() (T, error) {
	if len(*q) == 0 {
		var zero T
		return zero, errors.New("QueueEmpty")
	}
	v := (*q)[0]
	*q = (*q)[1:]
	return v, nil
}

func (q *Queue[T]) Empty() bool {
	return len(*q) == 0
}

func makeBinaryTreeFromArray(a []int) *TreeNode {
	if len(a) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: a[0],
	}
	q := Queue[*TreeNode]{root}
	indices := Queue[int]{0}
	for !q.Empty() {
		node, _ := q.Dequeue()
		index, _ := indices.Dequeue()

		leftNodeIndex := 2*index + 1
		rightNodeIndex := 2*index + 2

		if leftNodeIndex < len(a) {
			node.Left = &TreeNode{Val: a[leftNodeIndex]}
			q.Enqueue(node.Left)
			indices.Enqueue(leftNodeIndex)
		}

		if rightNodeIndex < len(a) {
			node.Right = &TreeNode{Val: a[rightNodeIndex]}
			q.Enqueue(node.Right)
			indices.Enqueue(rightNodeIndex)
		}
	}
	return root
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

func main() {
	root := []int{4, 2, 7, 1, 3, 6, 9}
	fmt.Println(invertTree(makeBinaryTreeFromArray(root)))
}
