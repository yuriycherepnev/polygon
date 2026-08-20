package main

import "fmt"

func main() {
	fmt.Println(remove([]int{}))
	fmt.Println(remove([]int{0}))
	fmt.Println(remove([]int{1, 0, 0, 2}))
}

func remove(in []int) []int {
	j := 0

	for i := 0; i < len(in); i++ {
		if in[i] != 0 {
			in[j] = in[i]
			j++
		}
	}
	return in[:j]
}
