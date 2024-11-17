package main

import (
	"fmt"
	"sort"
)

func unique(userIDs []int) []int {
	if len(userIDs) < 2 {
		return userIDs
	}

	sort.Slice(userIDs, func(i, j int) bool {
		return userIDs[i] < userIDs[j]
	})

	id := 0
	for i := 1; i < len(userIDs); i++ {
		if userIDs[id] != userIDs[i] {
			id++
			userIDs[id] = userIDs[i]
		}

	}
	return userIDs[:id+1]
}

func main() {
	slice := []int{1, 2, 6, 1, 6, 7, 4}
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	fmt.Println(slice)
	fmt.Println("no copy")
	newslice := make([]int, len(slice), cap(slice))
	newslice = unique(slice)
	fmt.Println(newslice)
}
