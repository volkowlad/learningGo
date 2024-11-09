package main

import (
	"fmt"
	"strings"
)

var invalidRequest = "invalid request"

type UserCreateRequest struct {
	FirstName string
	Age       int
}

func Validate(req UserCreateRequest) string {
	if req.FirstName == "" || strings.Contains(req.FirstName, " ") {
		return invalidRequest

	}

	if req.Age <= 0 || req.Age > 150 {
		return invalidRequest
	}

	return ""
}

func main() {
	test := UserCreateRequest{FirstName: "John", Age: 13}
	fnh := Validate(test)
	fmt.Println(fnh)
}
