package main

import "fmt"

func main() {
	fmt.Println(isMonotonic([]int{9, 5, 1}))
	fmt.Println(isMonotonic([]int{23, 5, 23}))
	fmt.Println(isMonotonic([]int{1, 5, 7, 8, 23}))
}

func isMonotonic(s []int) bool {
	isUp, isDown := true, true
	for i := 1; i < len(s); i++ {
		if s[i-1] >= s[i] {
			isUp = false
		}
		if s[i-1] <= s[i] {
			isDown = false
		}
	}
	return isUp || isDown
}
