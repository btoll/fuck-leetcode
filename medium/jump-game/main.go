package main

import "fmt"

//You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.
//
//Return true if you can reach the last index, or false otherwise.

func canJump(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	l := len(nums) - 1
	for i := 0; i < l; i++ {
		if l-i <= nums[i] {
			return true
		}
	}
	return false
}

func main() {
	//	nums := []int{2, 3, 1, 1, 4}
	nums := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums))
}
