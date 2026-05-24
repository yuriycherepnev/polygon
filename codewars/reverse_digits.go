package main

import "fmt"

func main() {
	count := Digitize(5555)
	fmt.Println(count)
}

func Digitize(number int) (slice []int) {
	for {
		slice = append(slice, number%10)
		number /= 10
		if number == 0 {
			return slice
		}
	}
}
