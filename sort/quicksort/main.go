package main

import "fmt"

func partition(nums []int, p, r int) int {
	i, j, pivot := p, p, nums[r]
	for i < r {
		if nums[i] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
		i++
	}
	nums[j], nums[r] = nums[r], nums[j]
	return j
}

func selection(nums []int, p, r int) {
	if p < r {
		q := partition(nums, p, r)
		selection(nums, p, q-1)
		selection(nums, q+1, r)
	}
}

func main() {
	nums := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Printf("nums=%#v\n", nums)
	selection(nums, 0, len(nums)-1)
	fmt.Printf("nums=%#v\n", nums)
}
