package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func GenerateSelfStory(name string, age int, money float64) string {
	return fmt.Sprintf("Hello! My name is %s. I'm %d y.o. And I also have $%g in my wallet right now.", name, age, money)
}

// в примере %.2f

func main() {
	p := Person{Name: "Vlad", Age: 23}
	fmt.Printf("%#v\n", p)
	name := "Vlad"
	age := 23
	money := 75000.34001
	res := GenerateSelfStory(name, age, money)
	fmt.Println(res)
}
