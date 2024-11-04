package main

import (
	"fmt"
	"strings"
)

func ModifySpaces(s, mode string) string {
	switch mode {
	case "dash":
		return strings.ReplaceAll(s, " ", "-")
	case "underscore":
		return strings.ReplaceAll(s, " ", "_")
	default:
		return strings.ReplaceAll(s, " ", "*")
	}
}

//func ModifySpaces(s, mode string) string {
//var replacement string

//switch mode {
//case "dash":
//replacement = "-"
//case "underscore":
//replacement = "_"
//default:
//replacement = "*"
//}

//return strings.ReplaceAll(s, " ", replacement)
//}

func main() {
	txt := " hello "
	md := "undersco"
	newtxt := ModifySpaces(txt, md)
	fmt.Println(newtxt)
}
