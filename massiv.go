package main

import (
	"fmt"
)

func SafeWrite(nums [5]int, i, val int) [5]int {
	if i < len(nums) && i >= 0 {
		nums[i] = val
	}
	return nums
}

func main() {
	msv := [5]int{1, 2, 57, 4, 5}
	i := 0
	val := 98
	msv = SafeWrite(msv, i, val)
	fmt.Println(msv)
}
