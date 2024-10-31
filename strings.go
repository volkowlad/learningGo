package main

import (
	"fmt"
	"strings"
)

func Greaatings(name string) string {
	name = strings.Trim(name, " ")
	name = strings.ToLower(name)
	return strings.Title(name)
}

func main() {
	nm := "iVan"
	nm = Greaatings(nm)
	fmt.Print(fmt.Sprintf("Hello, %s!", nm))
}
