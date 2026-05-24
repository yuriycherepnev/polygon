package main

import "fmt"

func main() {
	result := digitCount(6789)
	fmt.Println(result)
}

func digitCount(number int) int {
	count := 0
	for temp := number; temp != 0; temp /= 10 {
		count++
	}
	return count
}
