// Удалить все дубликаты из списка. In place
// Input: [ 3, 2, 1, 1, 0, 4, 5, 2, 0]
// Output: [ 3, 2, 1, 0, 4, 5]

package main

import "fmt"

func main() {
	singleNums := deleteDouble([]int{3, 2, 1, 1, 0, 4, 5, 2, 0})
	fmt.Println(singleNums)
}

func deleteDouble(numbers []int) []int {
	mapString := make(map[int]struct{})
	x := 0

	for index, value := range numbers {
		_, ok := mapString[value]
		if ok {
			continue
		}

		mapString[value] = struct{}{}
		numbers[x] = numbers[index]
		x++
	}

	return numbers[:x]
}
