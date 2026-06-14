package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	fields := strings.Fields(s)
	if len(fields) == 0 {
		return 0
	}
	return len(fields[len(fields)-1])
}

func main() {
	tests := []string{
		"Hello World",
		"   fly me   to   the moon  ",
		"luffy is still joyboy",
		"",
		" ",
	}
	for _, test := range tests {
		fmt.Println(lengthOfLastWord(test))
	}
}
