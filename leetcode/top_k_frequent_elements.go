package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1, 1}
	k := 2
	top := topKFrequent(nums, k)
	fmt.Println(top)
}

// bucket sort
func topKFrequent(nums []int, k int) []int {
	mapNums := make(map[int]int)

	for _, number := range nums {
		mapNums[number]++
	}

	frequentNums := make([][]int, len(nums)+1)

	for number, count := range mapNums {
		frequentNums[count] = append(frequentNums[count], number)
	}

	result := make([]int, 0, k)

	for i := len(frequentNums) - 1; i >= 0; i-- {
		for _, value := range frequentNums[i] {
			result = append(result, value)
			if len(result) == k {
				return result
			}
		}
	}

	return result
}

// решение через мапу
func topKFrequentMap(nums []int, k int) []int {
	mapNums := make(map[int]int)

	for _, v := range nums {
		mapNums[v]++
	}

	result := make([]int, k)

	i := 0
	for k > 0 {
		maxValue := 0
		maxKey := 0
		for key, value := range mapNums {
			if value > maxValue {
				maxValue = value
				maxKey = key
			}
		}
		result[i] = maxKey
		delete(mapNums, maxKey)
		k--
		i++
	}

	return result
}
