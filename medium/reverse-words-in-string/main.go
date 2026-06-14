package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	_ = strings.Fields(s)
	i := 0
	// Skip all spaces before the first word.
	for i < len(s) && s[i] == ' ' {
		i++
	}
	words := []string{}
	l := len(s)
	wordStart := i
	for i < l {
		for i < l && s[i] != ' ' {
			i++
			continue
		}
		words = append(words, s[wordStart:i])
		i++
		// Skip all spaces between words.
		for i < l && s[i] == ' ' {
			i++
			continue
		}
		wordStart = i
	}
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	tests := []string{
		"the sky is blue",
		"  hello world  ",
		"a good   example",
		"   a good  foo ",
		"   fffff ff gg ee ",
	}
	for _, test := range tests {
		fmt.Println(reverseWords(test))
	}
}
