package main

import "fmt"

func bubble(nums []int) {
	if len(nums) < 2 {
		return
	}
	l := len(nums)
	for l > 0 {
		for j := 1; j < l; j++ {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
		l--
	}
}

func main() {
	nums := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Printf("nums=%#v\n", nums)
	bubble(nums)
	fmt.Printf("nums=%#v\n", nums)
}
