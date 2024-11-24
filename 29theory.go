package main

import (
	"fmt"
)

type Printer interface {
	Print()
}

type UserEr struct {
	email string
}

func (u *UserEr) Print() {
	fmt.Println("My email is", u.email)
}

func TestPrint(p Printer) {
	p.Print()
}

func main() {
	u := &UserEr{email: "test"}
	u.Print()
	TestPrint(u)
}
