package main

import (
	"fmt"
)

func Shift(s string, step int) string {
	res := ""
	for i := 0; i < len(s); i++ {
		news := byte(int(s[i]) + step)
		res += string(news)
	}
	return res
}

func SH(s string, step int) string {
	if step == 0 {
		return s
	}

	shifted := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		shift := step + int(s[i])
		shifted[i] = byte(shift)
	}

	return string(shifted)
}

func main() {
	s := "abc"
	ns := Shift(s, -1)
	fmt.Println(ns)
}
