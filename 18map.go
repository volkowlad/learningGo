package main

import (
	"fmt"
)

func uniquenew(userID []int) []int {

	processed := make(map[int]struct{}) //создаем пустую мапу
	uniqueID := make([]int, 0)          // создаем пустой масив

	for _, uid := range userID { // смотрим массив по значению
		_, ok := processed[uid] // смотрим есть ли ключ в мапе
		if ok {                 // если есть, то продолжаем
			continue
		} // если нет, то

		uniqueID = append(uniqueID, uid) // добавляем к концу созданого массива значение, которого нет в ключе
		processed[uid] = struct{}{}      // создаем ключ с добавленным в массив значением с пустой ячейкой

	}

	return uniqueID // возвращаем полученный массив
}

func main() {
	cache := make(map[string]struct{})
	fmt.Println(cache)
	cache["key"] = struct{}{}
	_, ok := cache["key"]
	fmt.Println(cache, ok)
	mas := []int{55, 55, 3, 2, 55}
	mas = uniquenew(mas)
	fmt.Println(mas)
}
