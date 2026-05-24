package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(stringProcessing("reverse word aloha"))
	fmt.Println(stringProcessing("Hey fellow warriors")) // "Hey wollef sroirraw"
	fmt.Println(stringProcessing("This is a test"))      // "This is a test"
	fmt.Println(stringProcessing("This is another test"))
	fmt.Println(stringProcessing("Привет гайз , как делишки сегодня?"))
}

func stringProcessing(str string) string {
	words := strings.Split(str, " ")
	for i, word := range words {
		if len([]rune(word)) >= 5 {
			words[i] = reverseString([]byte(word))
		}
	}
	return strings.Join(words, " ")
}

/*
метод из chatGpt
*/
func reverseWord(word string) string {
	runes := []rune(word)
	for i, j := 0, len([]rune(word))-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

/*
метод glonas
*/
func reverseString(str []byte) string {
	index1 := 0
	index2 := len(str) - 1
	fmt.Println(str)

	for index1 < index2 {
		str[index1], str[index2] = str[index2], str[index1]
		index1++
		index2--
	}
	return string(str)
}
