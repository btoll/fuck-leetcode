package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	alpha := [26]rune{}
	for _, v := range s {
		alpha[v-'a'] += 1
	}
	for _, v := range t {
		alpha[v-'a'] -= 1
		if alpha[v-'a'] < 0 {
			return false
		}
	}
	for _, v := range alpha {
		if v != 0 {
			return false
		}
	}
	return true
}

// For Unicode:
//func isAnagram(s string, t string) bool {
//	if len(s) != len(t) {
//		return false
//	}
//	m := make(map[rune]int)
//	for _, v := range s {
//		m[v] += 1
//	}
//	for _, v := range t {
//		m[v] -= 1
//		if m[v] < 0 {
//			return false
//		}
//	}
//	for _, v := range m {
//		if v != 0 {
//			return false
//		}
//	}
//	return true
//}

func main() {
	//	s := "anagram"
	//	t := "nagaram"
	s := "rat"
	t := "car"
	fmt.Println(isAnagram(s, t))
}
