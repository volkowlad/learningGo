package main

import (
	"fmt"
	"strings"
	"unicode"
)

func LatinLettersMy(s string) string {
	res := make([]rune, 0)
	for _, char := range s {
		if unicode.Is(unicode.Latin, char) {
			res = append(res, char)
		}
	}
	return string(res)
}

func LatinLetters(s string) string {
	sb := strings.Builder{}
	for _, char := range s {
		if unicode.Is(unicode.Latin, char) {
			sb.WriteRune(char)
		}
	}
	return sb.String()
}

func main() {
	s := "bot привет world!"
	fmt.Println(LatinLettersMy(s))
	fmt.Println(len(LatinLettersMy(s)))
	fmt.Println(LatinLetters(s))
	fmt.Println(len(LatinLetters(s)))
}
