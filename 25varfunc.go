package main

import (
	"fmt"
)

func Merge(list ...[]int) []int {
	res := []int{}
	for i, _ := range list {
		res = append(res, list[i]...)
	}
	return res
}

func MergeHexlet(list ...[]int) []int {
	cap := 0
	for i := 0; i < len(list); i++ {
		cap += len(list[i])
	}

	res := make([]int, 0, cap)

	for _, n := range list {
		res = append(res, n...)
	}

	return res

}

func main() {
	list1 := []int{1, 2}
	list2 := []int{3}
	list3 := []int{4}
	res1 := Merge(list1, list2, list3)
	res2 := MergeHexlet(list1, list2, list3)
	fmt.Println(res1, res2)
}
