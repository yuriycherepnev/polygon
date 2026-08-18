// метод генерации массива случайных неповторяющихся чисел

package main

import (
	"fmt"
	"math/rand"
)

// rand.Intn(100) 0 <= n < 100
// для диапазона 50 - 100
// rand.Intn(51) + 50

func main() {
	uniqueNumbers := generateUniqueNumbersMap(5, 45, 50)
	fmt.Println(uniqueNumbers)

	uniqueNumbers2 := generateUniqueNumbersSlice(5, 45, 50)
	fmt.Println(uniqueNumbers2)
}

func generateUniqueNumbersMap(n, min int, max int) []int {
	if n > max {
		return nil
	}
	nums := make(map[int]struct{})

	result := make([]int, 0, n)
	for len(result) < n {
		randNum := rand.Intn(max-min) + min

		_, ok := nums[randNum]
		if ok {
			continue
		}
		nums[randNum] = struct{}{}
		result = append(result, randNum)
	}
	return result
}

func generateUniqueNumbersSlice(n, min int, max int) []int {
	if n > max {
		return nil
	}
	nums := make([]int, max)

	for i := min; i < max; i++ {
		nums[i] = i
	}
	rangeNums := nums[min:max]

	rand.Shuffle(len(rangeNums), func(i, j int) {
		rangeNums[i], rangeNums[j] = rangeNums[j], rangeNums[i]
	})

	return rangeNums[:n]
}
