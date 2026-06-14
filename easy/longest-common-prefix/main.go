package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	minLen := len(strs[0])
	for _, str := range strs[1:] {
		if l := len(str); l < minLen {
			minLen = l
		}
	}
	var pos int
	for pos = 0; pos < minLen; pos++ {
		ref := strs[0][pos]
		for _, str := range strs[1:] {
			if str[pos] != ref {
				return strs[0][:pos]
			}
		}
	}
	return strs[0][:minLen]
}

func main() {
	tests := [][]string{
		{"flower", "flow", "flight"},
		{"dog", "racecar", "car"},
		{},
		{"single"},
	}
	for _, t := range tests {
		fmt.Printf("%v -> %q\n", t, longestCommonPrefix(t))
	}
}
