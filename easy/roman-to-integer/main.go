package main

import (
	"fmt"
	"strings"
)

var roman = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	s = strings.ToUpper(s)
	var d int
	for i := 0; i < len(s); i++ {
		if (s[i] == 'I' || s[i] == 'X' || s[i] == 'C') &&
			i < len(s)-1 {
			// the key insight is that the conditional below should ONLY pass
			// if the current rune (s[i]) is less than the next one. if you
			// think about it, 100 (C) can only be used to subtract a number
			// reater than itself.  this principle applies to every number.
			if roman[rune(s[i+1])] > roman[rune(s[i])] {
				d += roman[rune(s[i+1])] - roman[rune(s[i])]
				i += 1
				continue
			}
		}
		d += roman[rune(s[i])]
	}
	return d
}

func main() {
	tests := []string{
		"xxix",
		"xxx",
		"v",
		"xiv",
		"ii",
		"im",
		"mcmxciv",
		"MDCCL",
	}

	for _, test := range tests {
		fmt.Println(romanToInt(test))
	}
}
