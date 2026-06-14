package main

import "fmt"

var roman = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func decimalLength(num int) int {
	var k int
	for num > 0 {
		num /= 10
		k += 1
	}
	return k
}

func intToRoman(num int) string {
	fmt.Printf("num=%+v\n", num)
	k := 0
	for num > 0 {
		if k%2 == 0 {
			num = num / 2
			fmt.Printf("num=%+v\n", num)
		} else {
			num = num / 5
			fmt.Printf("num=%+v\n", num)
		}
		k += 1
	}
	return ""
}

func main() {
	tests := []int{
		51,
		3,
		100,
		1750,
	}
	// MDCCL
	for _, test := range tests {
		fmt.Println(intToRoman(test))
		fmt.Println(decimalLength(test))
	}
}
