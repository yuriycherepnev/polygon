package main

import "fmt"

func main() {
	number := 56784

	result := numberToSlice(number)
	fmt.Println(result)
}

func numberToSlice(number int) []int {
	if number == 0 {
		return []int{0}
	}
	count := 0

	for temp := number; temp != 0; temp /= 10 {
		count++
	}

	numberSlice := make([]int, count)

	for i := count - 1; i >= 0; i-- {
		remainder := number % 10
		number /= 10
		numberSlice[i] = remainder
	}

	return numberSlice
}
