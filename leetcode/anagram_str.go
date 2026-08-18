// Определить, являются ли две строки анаграммами. Строка А является анаграммой строки Б,
// если можно переставить местами символы в строке А и получить строку Б.
// Input: a="лапоть", b="пальто"
// Output: true

package main

import "fmt"

func main() {
	singleNums := anagramStr("пальто", "лапоть")
	fmt.Println(singleNums)
}

func anagramStr(strOne string, strTwo string) bool {
	mapString := make(map[int32]int)

	for _, value := range strOne {
		mapString[value]++
	}
	for _, value := range strTwo {
		mapString[value]--
	}
	for _, value := range mapString {
		if value > 0 {
			return false
		}
	}
	return true
}
