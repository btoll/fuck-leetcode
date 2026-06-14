package main

import "fmt"

// Given an integer array nums and an integer k, return true if there are two distinct indices i and j in the array such that nums[i] == nums[j] and abs(i - j) <= k.

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i, v := range nums {
		if index, found := m[v]; found {
			if i-index <= k {
				return true
			}
		}
		m[v] = i
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 1}
	k := 3
	//	nums := []int{1, 0, 1, 1}
	//	k := 1
	//	nums := []int{1, 2, 3, 1, 2, 3}
	//	k := 2
	fmt.Println(containsNearbyDuplicate(nums, k))
}
