package main

import (
	"fmt"
)

func PopWord(words []string) string {
	m := make(map[string]int)
	cache := make(map[string]struct{})
	for _, word := range words {
		_, ok := cache[word]
		if ok {
			m[word]++
		} else {
			m[word] = 1
			cache[word] = struct{}{}
		}
	}
	resword := words[0]
	resnum := m[words[0]]
	for wr, nm := range m {
		if nm > resnum {
			resword = wr
		}
	}
	return resword
}

func main() {
	test := make(map[string]int)
	fmt.Println(test)
	s := "darov"
	test[s] = 1
	fmt.Println(test)
	words := []string{"darov", "darov", "bey"}
	test1 := PopWord(words)
	fmt.Println(test1)
}
