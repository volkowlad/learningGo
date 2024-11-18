package main

import (
	"fmt"
)

func Next(b byte) byte {
	return b + 1
}

func Prev(b byte) byte {
	return b - 1
}

func main() {
	a := byte('a')
	a += 1
	fmt.Println(a)
	Na := Next(a)
	Pa := Prev(a)
	fmt.Println(Na, Pa)
}
