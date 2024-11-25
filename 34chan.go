package main

import (
	"fmt"
)

func Work(numCh chan []int, sumCh chan int) {
	for num := range numCh {
		sumCh <- sum(num)
	}
}

func sum(num []int) int {
	s := 0
	for _, nm := range num {
		s += nm
	}
	return s
}

func main() {
	numCh := make(chan []int)
	sumCh := make(chan int)

	go Work(numCh, sumCh)

	numCh <- []int{10, 10}
	res := <-sumCh
	fmt.Println(res)
}
