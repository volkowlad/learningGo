package main

import (
	"fmt"
	"strconv"
)

//var test int
//var fnh string

func cnv(str int) string {
	return strconv.Itoa(str)
}

func main1() {

	test := -8
	fnh := cnv(test)
	fmt.Print(fnh)

}
