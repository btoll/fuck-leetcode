package main

import "fmt"

func siftDown(nums []int, index, size int) {
	i := index
	for i < size {
		child := 2*i + 1
		smaller := -1
		if child < size {
			smaller = child
			if child+1 < size && nums[child] < nums[child+1] {
				smaller = child + 1
			}
		}
		if smaller == -1 || nums[i] > nums[smaller] {
			break
		}
		nums[i], nums[smaller] = nums[smaller], nums[i]
		i = smaller
	}
}

func heapsort(nums []int) {
	l := len(nums)
	if l < 2 {
		return
	}
	for i := (l - 1) / 2; i >= 0; i-- {
		siftDown(nums, i, l)
	}
	size := l - 1
	for range nums {
		nums[0], nums[size] = nums[size], nums[0]
		siftDown(nums, 0, size)
		size--
	}
}

func main() {
	nums := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	fmt.Printf("nums=%#v\n", nums)
	heapsort(nums)
	fmt.Printf("nums=%#v\n", nums)
}
