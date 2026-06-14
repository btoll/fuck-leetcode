package main

import "fmt"

//func rotate(nums []int, k int) {
//	a := make([]int, len(nums))
//	l := len(nums)
//	for i, v := range nums {
//		n := (i + k) % l
//		a[n] = v
//	}
//	copy(nums, a)
//}

func work(nums []int, p, r int) {
	for i, j := p, r-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func rotate(nums []int, k int) {
	if k > len(nums) {
		return
	}
	work(nums, 0, len(nums))
	work(nums, 0, k)
	work(nums, k, len(nums))
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("nums=%+v\n", nums)
	rotate(nums, 3)
	fmt.Printf("nums=%+v\n", nums)
}
