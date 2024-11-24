package main

import (
	"fmt"
	"math"
)

type Dog struct {
	Is bool
}

type Counter struct {
	Value int
}

func (d Dog) Bark() {
	d.Is = true
}

func (a *Dog) UBark() {
	a.Is = true
}

func (c *Counter) Inc(delta int) {
	if delta == 0 {
		c.Value += 1
	}

	c.Value += delta
}

func (c *Counter) Dec(delta int) {
	if delta == 0 {
		c.Value -= 1
	}

	c.Value -= delta

	if c.Value < 0 {
		c.Value = 0
	}

}

func (c *Counter) DecH(delta int) {
	if delta == 0 {
		delta = 1
	}
	c.Value = int(math.Max(float64(c.Value-delta), float64(0)))
}

func main() {
	d := Dog{}
	d.Bark()
	a := Dog{}
	a.UBark()
	c := &Counter{}
	c.Value = 32
	c.Inc(5)
	c.Dec(100)
	c.Inc(32)
	c.Dec(0)
	fmt.Println(d.Is, a.Is)
	fmt.Println(c.Value)
	c.DecH(100)
	fmt.Printf("Hexlet function %d", c.Value)
	fmt.Println()
}
