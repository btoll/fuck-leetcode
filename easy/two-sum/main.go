package main

import "fmt"

//func twoSum(nums []int, target int) []int {
//	m := make(map[int]int)
//	for i, v := range nums {
//		m[v] = i
//	}
//	for i, v := range nums {
//		if index, found := m[target-v]; found && i != index {
//			return []int{i, index}
//		}
//	}
//	return nil
//}

func bubble(nums []int) {
	l := len(nums)
	for l > 0 {
		for i := 1; i < l; i++ {
			if nums[i-1] > nums[i] {
				nums[i-1], nums[i] = nums[i], nums[i-1]
			}
		}
		l -= 1
	}
}

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

func twoSum(nums []int, target int) []int {
	var candidate int
	for i, j := 0, len(nums)-1; i < j; {
		candidate = nums[i] + nums[j]
		if target == candidate {
			return []int{i, j}
		}
		if candidate > target {
			j -= 1
		} else {
			i += 1
		}
	}
	return nil
}

func main() {
	//	nums := []int{2, 7, 11, 15}
	//	nums := []int{3, 2, 4}
	nums := []int{3, 3}
	bubble(nums)
	fmt.Printf("nums=%+v\n", nums)
	fmt.Println(twoSum(nums, 6))
}
