package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	seen := make(map[rune]int)
	windowStart, maxWindow := 0, 0
	for i, v := range s {
		if index, found := seen[v]; found {
			windowStart = max(windowStart, index+1)
		}
		seen[v] = i
		maxWindow = max(maxWindow, i-windowStart+1)
	}
	return maxWindow
}

func main() {
	s := "abcabcbb"
	//	s := "pwwkew"
	//	s := "tmmzuxt"
	fmt.Println(lengthOfLongestSubstring(s))
}
