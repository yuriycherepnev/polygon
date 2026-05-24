package main

import "fmt"

func main() {
	number := 10
	result := isPalindromeTwo(number)
	fmt.Println(result)
}

// сравнение всего числа
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	divider := 1

	for temp := x / 10; temp != 0; temp /= 10 {
		divider *= 10
	}

	for x > 0 {
		first := x / divider
		last := x % 10

		fmt.Println(first, last)

		if first != last {
			return false
		}

		x = (x % divider) / 10
		divider /= 100
	}

	return true
}

// сравнение половины числа
func isPalindromeTwo(x int) bool {
	if x == 0 {
		return true
	}
	if x < 0 || x%10 == 0 {
		return false
	}
	reversed := 0

	for x > reversed {
		reversed = reversed*10 + x%10
		x /= 10
	}

	return x == reversed || x == reversed/10
}
