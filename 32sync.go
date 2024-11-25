package main

import (
	"fmt"
)

func MaxSum(num1, num2 []int) []int {
	sum1 := 0
	sum2 := 0
	for _, nmb1 := range num1 {
		sum1 += nmb1
	}

	for _, nmb2 := range num2 {
		sum2 += nmb2
	}
	if sum1 < sum2 {
		return num2
	}
	return num1
}

func main() {
	res1 := MaxSum([]int{1, 2, 3}, []int{10, 20, 30})
	res2 := MaxSum([]int{1, 2, 3}, []int{3, 2, 1})
	fmt.Println(res1)
	fmt.Println(res2)
}
