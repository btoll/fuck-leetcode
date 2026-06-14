package main

import "fmt"

type Stack[T comparable] []T

func NewStack[T comparable]() *Stack[T] {
	return new(Stack[T])
}

func (s *Stack[T]) Len() int {
	return len(*s)
}

func (s *Stack[T]) Pop() T {
	if len(*s) == 0 {
		panic("stack is empty")
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s *Stack[T]) Push(v T) {
	*s = append(*s, v)
}

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
// An input string is valid if:
//
// 1. Open brackets must be closed by the same type of brackets.
// 2. Open brackets must be closed in the correct order.
// 3. Every close bracket has a corresponding open bracket of the same type.
func isValid(s string) bool {
	if len(s) < 2 {
		return false
	}
	t := NewStack[rune]()
	brackets := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	for _, bracket := range s {
		if closing, found := brackets[bracket]; found {
			t.Push(closing)
		} else {
			if t.Len() == 0 || bracket != t.Pop() {
				return false
			}
		}
	}
	if t.Len() > 0 {
		return false
	}
	return true
}

func main() {
	s := "()[]{}"
	//	s := "){"
	fmt.Println(isValid(s))
}
