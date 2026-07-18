package main

import "fmt"

func main() {
	var a []int
	for i := range 3 {
		a = append(a, i+1)
		fmt.Println(a, len(a), cap(a))
	}

	b := append(a, 4)
	c := append(a, 5)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
