package main

import (
	"fmt"
)

func RemoveFirstElement(slice []int) []int {
	if len(slice) == 0 {
		return slice
	} else {
		return slice[1:]
	}
}

func main() {
	ogn := []int{1, 2, 3, 4, 5}
	ogn1 := make([]int, 0, 5)
	fmt.Println(RemoveFirstElement(ogn))
	fmt.Println(RemoveFirstElement(ogn1))
}
