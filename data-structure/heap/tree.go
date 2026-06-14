package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func makeCompleteTree(nums []int) *TreeNode {
	root := &TreeNode{Val: nums[0]}
	q := []*TreeNode{root}
	i := 0
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		left := 2*i + 1
		right := 2*i + 2
		if left < len(nums) {
			node.Left = &TreeNode{Val: nums[left]}
			q = append(q, node.Left)
		}
		if right < len(nums) {
			node.Right = &TreeNode{Val: nums[right]}
			q = append(q, node.Right)
		}
		i++
	}
	return root
}

func bfs(root *TreeNode) []int {
	q := []*TreeNode{root}
	nums := []int{}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		nums = append(nums, node.Val)
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}
	return nums
}
