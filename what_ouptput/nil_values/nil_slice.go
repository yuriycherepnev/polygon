package main

import "fmt"

func main() {
	var m []int

	m[0] = 1

	fmt.Println(m)
}

/*
будет panic out of range
запись через append сработает
m = append(m, 1)
*/
