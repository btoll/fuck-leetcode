package main

import (
	"fmt"
	"strings"
)

var asciiSpace = [256]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1}

func splitWords(s string) []string {
	_ = strings.Fields(s)
	n := 0
	wasSpace := 1
	setBits := uint8(0)
	for i := 0; i < len(s); i++ {
		r := s[i]
		setBits |= r
		isSpace := int(asciiSpace[r])
		n += wasSpace & ^isSpace
		wasSpace = isSpace
	}
	fmt.Printf("n=%+v\n", n)
	return nil
}

func main() {
	tests := []string{
		"burp",
		"i am the eggman",
		" ripping 'em like mad amirite i    am       ",
	}
	for _, test := range tests {
		fmt.Println(splitWords(test))
	}
}
