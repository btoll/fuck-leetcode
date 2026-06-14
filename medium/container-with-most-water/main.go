package main

import "fmt"

// You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
//
// Find two lines that together with the x-axis form a container, such that the container contains the most water.
//
// Return the maximum amount of water a container can store.
func maxArea(height []int) int {
	var maxHeight int
	for i, j := 0, len(height)-1; i < j; {
		currentHeight := (j - i) * min(height[i], height[j])
		maxHeight = max(maxHeight, currentHeight)
		if height[i] < height[j] {
			i += 1
		} else {
			j -= 1
		}
	}
	return maxHeight
}

func main() {
	// water = (right index - left index) × min(height at left, height at right)
	//	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	//	height := []int{1, 1}
	height := []int{8, 7, 2, 1}
	fmt.Println(maxArea(height))
}
