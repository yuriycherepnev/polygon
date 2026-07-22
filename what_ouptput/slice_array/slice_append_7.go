package main

import "fmt"

func main() {
	s := make([]int, 0, 5) // [], 0 0 0 0 0, len 0 cap 5
	s = append(s, 1, 2, 3) //[1 2 3], 0 0, len 3 cap 5

	subSlice := s[1:3] // [2 3] 0 0, len 2 cap 4

	subSlice[0] = 99               // [99 3] 0 0, len 2 cap 4
	subSlice = append(subSlice, 4) // [99 3 4] 0 , len 3 cap 4

	s = append(s, 5, 6, 7) // [1 99 3 5 6 7] 0 0 0 0, len 6 cap 10, здесь создан новый массив

	subSlice[1] = 100 // [99 100 4] 0

	fmt.Println("s:", s)               // [1 99 3 5 6 7] 0 0 0 0, len 6 cap 10
	fmt.Println("subSlice:", subSlice) // 0 [99 100 4] 0, len 3 cap 4
}
