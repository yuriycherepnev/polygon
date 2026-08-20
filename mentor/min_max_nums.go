// Дан массив чисел, содержащий минимум два элемента. Нужно найти максимальное произведение двух элементов в этом массиве.
// Input: nums = [-2, 1, -3, 4, -1, 2, 1, -5, 4]
// Output: 16

package main

import "fmt"

func main() {
	multiplierNumber := nums([]int{-2, 1, -3, 4, -1, 2, 1, -5, -5, 4})
	fmt.Println(multiplierNumber)
}

func nums(numbers []int) int {
	maxNumber := [2]int{}
	minNumber := [2]int{numbers[0], numbers[1]}

	for _, value := range numbers {
		if value > maxNumber[0] {
			maxNumber[0] = value
		}
		if maxNumber[1] < maxNumber[0] {
			maxNumber[0], maxNumber[1] = maxNumber[1], maxNumber[0]
		}
		if value < minNumber[0] {
			minNumber[0] = value
		}
		if minNumber[1] > minNumber[0] {
			minNumber[0], minNumber[1] = minNumber[1], minNumber[0]
		}
	}
	maxMultiply := maxNumber[0] * maxNumber[1]
	minMultiply := minNumber[0] * minNumber[1]
	if maxMultiply > minMultiply {
		return maxMultiply
	}
	return minMultiply
}
