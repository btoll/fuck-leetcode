package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern, s string) bool {
	words := strings.Fields(s)
	if len(pattern) != len(words) {
		return false
	}
	m := make(map[rune]string)
	rev := make(map[string]rune)
	for i, c := range pattern {
		w := words[i]
		if val, found := m[c]; found {
			if val != w {
				return false
			}
		} else {
			if rc, found := rev[w]; found && rc != c {
				return false
			}
			m[c] = w
			rev[w] = c
		}
	}
	return true
}

func main() {
	tests := []struct {
		pattern string
		s       string
	}{
		{pattern: "abba", s: "dog cat cat dog"},
		{pattern: "abba", s: "dog cat cat fish"},
		{pattern: "aaaa", s: "dog cat cat dog"},
		{pattern: "bbbb", s: "cat cat cat cat"},
		{pattern: "abba", s: "dog dog dog dog"},
	}
	for _, test := range tests {
		fmt.Println(wordPattern(test.pattern, test.s))
	}
}
