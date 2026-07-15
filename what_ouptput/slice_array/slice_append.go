package main

import "fmt"

func main() {
	var a []int // [] len 0 cap 0

	fmt.Println(len(a), cap(a))
	for i := range 3 {
		a = append(a, i+1)

		fmt.Println(len(a), cap(a))
	}
}
