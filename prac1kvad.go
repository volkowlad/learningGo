package main

import "fmt"

func Kvad(num int, kvadCh chan int) {
	kvadCh <- num * num
}

func SumKvad(nums []int) int {
	resCh := make(chan int)
	sum := 0

	for _, numi := range nums {
		go Kvad(numi, resCh)
		//sum += <- resCh
	}

	for range nums {
		sum += <-resCh
	}

	return sum
}

func KvadTea(nums []int) int {
	kvadCh := make(chan int)
	s := 0
	for _, ni := range nums {
		go func(n int) {
			kvadCh <- n * n
		}(ni)
	}
	for range nums {
		s += <-kvadCh
	}
	return s

}

func main() {
	test := []int{1, 2, 3}
	res := SumKvad(test)
	res2 := KvadTea(test)
	fmt.Println(res, res2)
}
