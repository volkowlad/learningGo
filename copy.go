package main

import (
	"fmt"
)

func IntsCopy(src []int, maxLen int) []int {
	if maxLen <= 0 {
		return make([]int, 0)
	} else if maxLen > len(src) {
		cp := make([]int, len(src))
		copy(cp, src)
		return cp
	} else {
		cp := make([]int, maxLen)
		copy(cp, src)
		return cp
	}
}

func main() {
	src := []int{1, 2, 3}
	len := 1
	copysrc := IntsCopy(src, len)
	fmt.Println(copysrc)
}
