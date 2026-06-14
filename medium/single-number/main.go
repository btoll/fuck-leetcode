package main

import "fmt"

//Given an integer array nums where every element appears three times except for one, which appears exactly once. Find the single element and return it.
//You must implement a solution with a linear runtime complexity and use only constant extra space.

func singleNumber(nums []int) int {
	var ones int
	var twos int

	for _, num := range nums {
		twos |= ones & num
		ones ^= num
		threes := ones & twos
		ones &= ^threes
		twos &= ^threes
	}

	return ones
}

func main() {
	//	nums := []int{2, 2, 3, 2}
	//	nums := []int{0, 1, 0, 1, 0, 1, 99}
	nums := []int{30000, 500, 100, 30000, 100, 30000, 100}
	fmt.Println(singleNumber(nums))
}
