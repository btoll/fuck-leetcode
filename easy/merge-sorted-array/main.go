package main

import "fmt"

//You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.
//
//Merge nums1 and nums2 into a single array sorted in non-decreasing order.
//
//The final sorted array should not be returned by the function, but instead be stored inside the array nums1. To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := 0, 0, 0
	left := make([]int, m)
	right := make([]int, n)

	copy(left, nums1[:m])
	copy(right, nums2)

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			nums1[k] = left[i]
			i += 1
		} else {
			nums1[k] = right[j]
			j += 1
		}
		k += 1
	}

	for i < len(left) {
		nums1[k] = left[i]
		i += 1
		k += 1
	}

	for j < len(right) {
		nums1[k] = right[j]
		j += 1
		k += 1
	}
}

func main() {
	// nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	fmt.Printf("nums1=%+v\n", nums1)
	merge(nums1, 3, nums2, len(nums2))
	fmt.Printf("nums1=%+v\n", nums1)
}
