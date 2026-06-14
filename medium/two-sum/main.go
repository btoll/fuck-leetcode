package main

import "fmt"

// Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two numbers such that they add up to a specific target number. Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length.
//
// Return the indices of the two numbers index1 and index2, each incremented by one, as an integer array [index1, index2] of length 2.
//
// The tests are generated such that there is exactly one solution. You may not use the same element twice.
//
// Your solution must use only constant extra space.
//func twoSum(numbers []int, target int) []int {
//	m := make(map[int]int)
//	for i, v := range numbers {
//		m[v] = i
//	}
//	for i, v := range numbers {
//		if index, found := m[target-v]; found && index != i {
//			return []int{i + 1, index + 1}
//		}
//	}
//	return nil
//}

func twoSum(numbers []int, target int) []int {
	for i, j := 0, len(numbers)-1; i < j; {
		candidate := numbers[i] + numbers[j]
		if candidate == target {
			return []int{i + 1, j + 1}
		}
		if candidate > target {
			j -= 1
		} else {
			i += 1
		}
	}
	return nil
}

//func twoSum(numbers []int, target int) bool {
//	m := make(map[int]int)
//	for i, v := range numbers {
//		m[v] = i
//	}
//	for i, v := range numbers {
//		if index, found := m[target-v]; found && index != i {
//			return true
//		}
//	}
//	return false
//}

func main() {
	numbers := []int{2, 7, 11, 15}
	fmt.Println(twoSum(numbers, 9))
}
