package main

import (
	"fmt"
	"unicode"
)

func IsASCII(s string) bool {
	for _, i := range s {
		if i > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func main() {
	s := "hello"
	fmt.Println(IsASCII(s))
}
