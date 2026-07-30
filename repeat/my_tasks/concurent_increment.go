package main

import (
	"fmt"
)

func increment(counter *int) {
	for i := 0; i < 100000; i++ {
		*counter++
	}
}

func main() {
	counter := 1

	go increment(&counter)
	go increment(&counter)

	fmt.Println(counter)
}
