package main

import (
	"array"
	"fmt"
)

func isValid(s string) bool {
	var newAS = array.NewDefaultArray()
	for _, c := range s {
		if c == '{' || c == '[' || c == '(' {
			newAS.Push(c)
			fmt.Println(newAS.ToString())
		} else {
			if newAS.IsEmpty() {
				return false
			}
			topChar := newAS.Pop()
			if c == '(' && rune(topChar) != ')' {
				return false
			}
			if c == '[' && topChar != ']' {
				return false
			}
			if c == '{' && topChar != '}' {
				return false
			}
		}

	}

	return newAS.IsEmpty()
}

func main() {

	char := "({}])"
	fmt.Println(isValid(char))
}
