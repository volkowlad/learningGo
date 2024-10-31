package main

import (
	"fmt"
	"math"
)

func minInt(x, y int) int {
	xflt := float64(x)
	yflt := float64(y)
	z := math.Min(xflt, yflt)
	return int(z)
}

func main() {
	x := -8
	y := 4
	fmt.Println(minInt(x, y))
}
