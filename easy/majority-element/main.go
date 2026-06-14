package main

import "fmt"

// When a majority always exists...
//func majorityElement(nums []int) int {
//	m := make(map[int]int)
//	threshold := len(nums) / 2
//	for _, v := range nums {
//		m[v]++
//		if m[v] > threshold {
//			return v
//		}
//	}
//	return -1
//}

// Boyer-Moore Voting Algorithm
// The candidate is initially set to the first element(why not -1?).
// Whenever the count reaches zero during iteration, the current element
// becomes the new candidate. The element that remains as the candidate
// after processing the entire array is the majority element.
func majorityElement(nums []int) int {
	count := 0
	candidate := -1
	for _, n := range nums {
		if count == 0 {
			candidate = n
		}
		if candidate == n {
			count += 1
		} else {
			count -= 1
		}
	}
	return candidate
}

func main() {
	//	nums := []int{2, 2, 1, 1, 1, 2, 2}
	//	nums := []int{2, 2, 1, 1, 1, 2, 2}
	nums := []int{3, 2, 3}
	fmt.Println("majority element", majorityElement(nums))
}
