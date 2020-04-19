package main

import (
	"array"
	"fmt"
	"strings"
)

func isValid(s string) bool {
	charSlice := strings.SplitAfter(s, "")
	var newAS = array.NewDefaultArray()
	for _, c := range charSlice {
		if c == "{" || c == "[" || c == "(" {
			newAS.Push(c)
			fmt.Println(newAS.ToString())
		} else {
			if newAS.IsEmpty() {
				return false
			}
			topChar := newAS.Pop()
			if c == "(" && topChar != ")" {
				return false
			}
			if c == "[" && topChar != "]" {
				return false
			}
			if c == "{" && topChar != "}" {
				return false
			}
		}

	}

	return newAS.IsEmpty()
}

func main() {

	char := "({})"
	fmt.Println(isValid(char))
}
