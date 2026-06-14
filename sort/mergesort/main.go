package main

import "fmt"

func merge(nums []int, p, q, r int) {
	i, j, k := 0, 0, p
	left := make([]int, q-p+1)
	right := make([]int, r-q)
	copy(left, nums[p:q+1])
	copy(right, nums[q+1:r+1])

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			nums[k] = left[i]
			i++
		} else {
			nums[k] = right[j]
			j++
		}
		k++
	}
	for i < len(left) {
		nums[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		nums[k] = right[j]
		j++
		k++
	}
}

func mergesort(nums []int, p, r int) {
	if p < r {
		q := (p + r) / 2
		mergesort(nums, p, q)
		mergesort(nums, q+1, r)
		merge(nums, p, q, r)
	}
}

func main() {
	nums := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Printf("nums=%#v\n", nums)
	mergesort(nums, 0, len(nums)-1)
	fmt.Printf("nums=%#v\n", nums)
}
