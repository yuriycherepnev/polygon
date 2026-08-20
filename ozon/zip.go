package main

import "fmt"

func main() {
	s1, s2 := []int{1, 2, 3}, []int{4, 5, 6, 7, 8}

	fmt.Println(zip(s1, s2)) // [[1 4] [2 5] [3 6]]
}

func zip(s1 []int, s2 []int) [][]int {
	minLen := len(s1)
	if minLen > len(s2) {
		minLen = len(s2)
	}
	zipSlice := make([][]int, minLen)
	for i := 0; i < minLen; i++ {
		zipSlice = append(zipSlice, []int{s1[i], s2[i]})
	}
	return zipSlice
}

func zip2(s ...[]int) [][]int {
	minLen := len(s[0])
	for _, v := range s {
		if minLen > len(v) {
			minLen = len(v)
		}
	}
	zipSlice := make([][]int, minLen)

	for i := 0; i < minLen; i++ {
		x := make([]int, minLen)
		for l := 0; l < len(s); l++ {
			x = append(x, s[l][i])
		}
		zipSlice = append(zipSlice, x)
	}
	return zipSlice
}
