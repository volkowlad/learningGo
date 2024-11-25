package main

import (
	"fmt"
	"time"
)

func Sum(num []int) int {
	s := 0
	for _, nmb := range num {
		s += nmb
	}
	return s
}

func ParSum(num []int, res *int) {
	*res = Sum(num)
}

func SumGor(num1, num2 []int) []int {
	s1, s2 := 0, 0

	go ParSum(num1, &s1)
	go ParSum(num2, &s2)

	time.Sleep(100 * time.Millisecond)

	if s2 < s1 {
		return num1
	}
	return num2
}

func main() {
	res1 := SumGor([]int{1, 2, 3}, []int{10, 20, 30})
	res2 := SumGor([]int{1, 2, 3}, []int{3, 2, 1})
	fmt.Println(res1)
	fmt.Println(res2)
}
