package main

import "fmt"

func strStr(haystack string, needle string) int {
	var window int
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[window] {
			window += 1
			if window == len(needle) {
				return i - len(needle) + 1
			}
		} else {
			if window > 0 {
				i -= window
			}
			window = 0
		}
	}
	return -1
}

func main() {
	//	haystack := "sadbutsad"
	//	needle := "but"
	haystack := "mississippi"
	needle := "issip"
	fmt.Println(strStr(haystack, needle))
}
