package main

import "fmt"

func selection(nums []int) {
	for i := 0; i < len(nums); i++ {
		minIndex := i
		for j := i; j < len(nums)-i; j++ {
			if nums[j] < nums[i] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
}

func main() {
	nums := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Printf("nums=%#v\n", nums)
	selection(nums)
	fmt.Printf("nums=%#v\n", nums)
}
