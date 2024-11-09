package main

import (
	"fmt"
)

const (
	OKmsg        = "OK"
	CANCELLEDmsg = "CANCELLED"
)

const (
	OK = iota
	CANCELLED
	UNKNOWN
)

func Error(msg string) int {
	switch msg {
	case OKmsg:
		return OK
	case CANCELLEDmsg:
		return CANCELLED
	default:
		return UNKNOWN
	}

}

func main() {
	msg := "O"
	out := Error(msg)
	fmt.Println(out)
}
