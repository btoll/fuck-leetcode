package main

import "fmt"

// Given a positive integer n, write a function that returns the number of in its binary representation (also known as the Hamming weight).
func hammingWeight(n int) int {
	if isPowerOf2(n) {
		return 1
	}
	var setBits int
	for n > 0 {
		if n&1 == 1 {
			setBits += 1
		}
		n >>= 1
	}
	return setBits
}

func isPowerOf2(n int) bool {
	return n > 0 && n&(n-1) == 0
}

func main() {
	fmt.Println(hammingWeight(128))
	// fmt.Println(isPowerOf2(n))
}
