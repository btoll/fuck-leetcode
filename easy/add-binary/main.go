package main

import (
	"fmt"
	"strings"
)

//func pow(base, exp int) int {
//	result := 1
//	for i := 0; i < exp; i++ {
//		result *= base
//	}
//	return result
//}

func steps(n int) int {
	var steps int
	for n > 0 {
		n >>= 1
		steps += 1
	}
	return steps
}

func intToString(n int) string {
	if n == 0 {
		return "0"
	}
	exp := steps(n)
	s := make([]string, exp)
	for pos := exp - 1; pos >= 0; pos-- {
		index := exp - 1 - pos
		if (n>>pos)&1 == 1 {
			s[index] = "1"
		} else {
			s[index] = "0"
		}
	}
	return strings.Join(s, "")
}

func stringToBinary(a string) int {
	var aN int
	for i, char := range a {
		if string(char) == "1" {
			shift := len(a) - i - 1
			aN += 1 << shift
		}
	}
	return aN
}

// Given two binary strings a and b, return their sum as a binary string.
func addBinary(a string, b string) string {
	return intToString(
		stringToBinary(a) + stringToBinary(b),
	)
}

func main() {
	a := "1010"
	b := "1011"
	//	a := "11"
	//	b := "1"
	//	a := "0"
	//	b := "0"
	fmt.Println(addBinary(a, b))
	fmt.Println(stringToBinary(a))
	fmt.Println(stringToBinary(b))
}
