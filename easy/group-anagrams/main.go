package main

import (
	"fmt"
	"sort"
	"strings"
)

//func bubble(word string) string {
//	sorted := strings.Split(word, "")
//	for i := range sorted {
//		for j := 1; j < len(sorted)-i; j++ {
//			if sorted[j-1] > sorted[j] {
//				sorted[j-1], sorted[j] = sorted[j], sorted[j-1]
//			}
//		}
//	}
//	return strings.Join(sorted, "")
//}

func groupAnagrams(words []string) [][]string {
	m := make(map[string][]string)
	for _, word := range words {
		toSort := strings.Split(word, "")
		sort.Strings(toSort)
		sorted := strings.Join(toSort, "")
		m[sorted] = append(m[sorted], word)
	}
	a := make([][]string, 0, len(m))
	for _, v := range m {
		a = append(a, v)
	}
	return a
}

func main() {
	words := []string{"eat", "tea", "ate", "bat", "tab", "cat"}
	fmt.Println(groupAnagrams(words))
}
