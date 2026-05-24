package main

import "fmt"

func main() {
	result := halfNumber(678876)

	fmt.Println(result)
}

func halfNumber(number int) int {
	half := 0
	multiply := 1

	for number > half {
		half += (number % 10) * multiply
		multiply *= 10
		number /= 10
	}

	return half
}
