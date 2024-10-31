package main

import (
	"fmt"
	"math"
)

//func minInt(x, y int) int {
//xflt := float64(x)
//yflt := float64(y)
//z := math.Min(xflt, yflt)
//return int(z)
//}

// solution on hexlet

func minInt(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func main() {
	x := -8
	y := 4
	fmt.Println(minInt(x, y))
}
