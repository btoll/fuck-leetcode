package main

import (
	"fmt"
)

//func reverseBits(n int) int {
//	var b strings.Builder
//	for n > 0 {
//		if n&1 == 1 {
//			b.WriteByte('1')
//		} else {
//			b.WriteByte('0')
//		}
//		n >>= 1
//	}
//	if b.Len() < 32 {
//		l := b.Len()
//		for l < 32 {
//			b.WriteByte('0')
//			l += 1
//		}
//	}
//	reversed, _ := strconv.ParseUint(b.String(), 2, 32)
//	return int(reversed)
//}

//func reverseBits(n int) int {
//	var k int
//	for v := 0; v < 32; v++ {
//		bit := (n >> v) & 1
//		k |= bit << (32 - 1 - v) // Use OR instead of addition
//	}
//	return k
//}

func reverseBits(n int) int {
	var k int
	for position := range 32 {
		bit := n >> position & 1
		k |= bit << (32 - position - 1)
	}
	return k
}

func main() {
	var n = 43261596
	//      964176192
	fmt.Println(reverseBits(n))
}
