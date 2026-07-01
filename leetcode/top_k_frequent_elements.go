package main

import "fmt"

func main() {
	nums := []int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}
	k := 2
	top := topKFrequent(nums, k)
	fmt.Println(top)
}

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
