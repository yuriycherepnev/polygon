package main

import "fmt"

func main() {
	number := 1236456
	result := reverseNumber(number)
	fmt.Println(result)
}

func reverseNumber(number int) int {
	reversed := 0
	for number > 0 {
		reversed = reversed*10 + number%10
		number /= 10
	}
	return reversed
}
