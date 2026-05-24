package main

import "fmt"

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func main() {
	arr := []int{2, 1, 3, 1, 2, 2, 2, 2, 4, 4, 4, 1}
	fmt.Println(PickPeaks(arr))
}

func PickPeaks(array []int) PosPeaks {
	posPeaks := PosPeaks{}
	candidate := 0
	for i := 1; i < len(array); i++ {
		if array[i-1] < array[i] {
			candidate = i
		} else if array[i-1] > array[i] && candidate > 0 {
			posPeaks.Pos = append(posPeaks.Pos, candidate)
			posPeaks.Peaks = append(posPeaks.Peaks, array[candidate])
			candidate = 0
		}
	}
	return posPeaks
}
