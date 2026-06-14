package main

import (
	"errors"
	"fmt"
)

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

func compare(leftNode, rightNode *TreeNode) bool {
	if leftNode == nil && rightNode == nil {
		return true
	}
	if leftNode == nil || rightNode == nil {
		return false
	}
	if leftNode.Val != rightNode.Val {
		return false
	}
	return compare(leftNode.Left, rightNode.Right) && compare(leftNode.Right, rightNode.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return compare(root.Left, root.Right)
}

func main() {
	root := []int{1, 2, 2, 3, 4, 4, 3}
	fmt.Println(isSymmetric(makeBinaryTreeFromArray(root)))
}
