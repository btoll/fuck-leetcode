package main

import "fmt"

//Given an integer array nums and an integer val, remove all occurrences of val in nums in-place. The order of the elements may be changed. Then return the number of elements in nums which are not equal to val.
//
//Consider the number of elements in nums which are not equal to val be k, to get accepted, you need to do the following things:
//
// - Change the array nums such that the first k elements of nums contain the elements which are not equal to val. The remaining elements of nums are not important as well as the size of nums.
// - Return k.

func removeElement(nums []int, val int) int {
	boundary := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if val == nums[i] {
			nums[i], nums[boundary] = nums[boundary], nums[i]
			boundary -= 1
		}
	}
	return boundary
}

func main() {
	//	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	nums := []int{3, 2, 2, 3}
	fmt.Printf("nums=%+v\n", nums)
	fmt.Println("k", removeElement(nums, 3))
	fmt.Printf("nums=%+v\n", nums)
}
