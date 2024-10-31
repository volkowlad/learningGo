package main

import (
	"fmt"
)

func IsValid(id int, text string) bool {
	return id > 0 && text != ""
}

func main() {
	id := -5
	text := "pet"
	fmt.Print(IsValid(id, text))
}
