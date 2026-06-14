package main

import "fmt"

// Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
//
// You must implement a solution with a linear runtime complexity and use only constant extra space.
func singleNumber(nums []int) int {
	v := nums[0]
	for i := 1; i < len(nums); i++ {
		v ^= nums[i]
	}
	return v
}

func main() {
	nums := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums))
}
