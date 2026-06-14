package main

import "fmt"

func xorSwap(a, b int) (int, int) {
	//	t := a ^ b
	//	return a ^ t, b ^ t
	return a ^ (a ^ b), b ^ (a ^ b)
}

func main() {
	// a, b := 5, 3
	// a, b := 55, 33
	a, b := 1001, 7777
	fmt.Println(xorSwap(a, b))
}
