package main

import (
	"fmt"
)

type counter int

func (c *counter) inc() {
	*c++
}

type Personuser struct {
	Age uint8
}

type PersonList []Personuser

func (pl PersonList) Get() map[uint8]int {
	res := make(map[uint8]int)
	for _, num := range pl {
		res[num.Age]++
	}
	return res
}

func main() {
	pl := PersonList{
		{Age: 18},
		{Age: 44},
		{Age: 18},
	}
	pop := pl.Get()
	fmt.Println(pop)
}
