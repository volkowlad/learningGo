package main

import (
	"fmt"
)

type User struct {
	email string
	pass  string
}

type Parent struct {
	Name      string
	Childdren []Child
}

type Child struct {
	Name string
	Age  int
}

func CopyPar(p *Parent) Parent {
	if p == nil {
		return Parent{}
	}

	children := make([]Child, len(p.Childdren))
	copy(children, p.Childdren)
	return Parent{
		Name:      p.Name,
		Childdren: children,
	}
}

func main() {
	u := &Parent{}
	fmt.Printf("try: %+v\n", u)
}
