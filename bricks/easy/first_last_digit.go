package main

import "fmt"

func main() {
	number := 57868765

	first := firstDigit(number)
	last := lastDigit(number)

	fmt.Println(first, last)
}

func firstDigit(number int) int {
	divider := 1
	for temp := number / 10; temp != 0; temp /= 10 {
		divider *= 10
	}
	return number / divider
}

func lastDigit(number int) int {
	return number % 10
}
