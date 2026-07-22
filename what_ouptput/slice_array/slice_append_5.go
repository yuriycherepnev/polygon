package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := a[2:4]
	c := append(b, 10)
	c[1] = 55

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
