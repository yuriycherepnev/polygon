package main

import "fmt"

const MAX = 5

func main() {
	s := generate() // s = [1,2,3,4], 0; len=4, cap=5
	mutation(s)     // s = [1,2,3,4], -1, len=4, cap=5
	fmt.Println("s=", s)
	fmt.Println(s[0:MAX])
}

func generate() []int {
	out := make([]int, 0, MAX) // out = [],0,0,0,0,0, len=0, cap=5
	for i := 1; i < MAX; i++ { // 1,2,3,4
		out = append(out, i)
	}
	return out
}

func mutation(s []int) {
	s = append(s, -1) // len=5, cap=5, s = [1,2,3,4, -1]
}
