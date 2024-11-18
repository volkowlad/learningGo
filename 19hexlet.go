package main

import (
	"fmt"
)

func HexletWord(words []string) string {
	wordsCount := make(map[string]int, 0) // создаем пустую мапу [строка]инт
	mostPopWord := ""                     // записываем пустое значение
	max := 0                              // записываем пустое значение

	for i := len(words) - 1; i >= 0; i -= 1 {
		word := words[i]             // достаем слово из массива
		wordsCount[word] += 1        // по ключу увеличиваем количесвто (создаем если ключ пустой)
		if wordsCount[word] >= max { // сравниваем значение ключа с максимальным
			max = wordsCount[word] // если оно больше, то максимальному присваиваем значение
			mostPopWord = word     // выбираем слово, которое мы достали с максимальным количеством
		}
	}

	return mostPopWord
}

func main() {
	words := []string{"darov", "darov", "bey"}
	test1 := HexletWord(words)
	fmt.Println(test1)
}
