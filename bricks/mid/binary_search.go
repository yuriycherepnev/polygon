package main

import "fmt"

func main() {
	array := []int{2, 4, 6, 8, 10, 1000, 1001, 1002, 2000}

	i, j := DoubleBinarySearch(array, 1002)
	fmt.Println(array[i], array[j])

	k := ClosestBinarySearch(array, 2002)
	fmt.Println(array[k])
}

func BinarySearch(arr []int, x int) int {
	i, j := 0, len(arr)-1
	for i < j {
		midIndex := int(uint(i+j)) >> 1
		if arr[midIndex] >= x {
			j = midIndex
		} else {
			i = midIndex + 1
		}
	}
	return i
}

func DownBinarySearch(arr []int, x int) int {
	i, j := 0, len(arr)
	for i+1 < j {
		midIndex := int(uint(i+j)) >> 1
		if arr[midIndex] > x {
			j = midIndex
		} else {
			i = midIndex
		}
	}
	return i
}

func DoubleBinarySearch(arr []int, x int) (int, int) {
	i, j := 0, len(arr)-1
	for i+1 < j {
		midIndex := int(uint(i+j)) >> 1
		if arr[midIndex] > x {
			j = midIndex
		} else {
			i = midIndex
		}
	}
	return i, j
}

func ClosestBinarySearch(arr []int, x int) int {
	i, j := 0, len(arr)-1
	for i+1 < j {
		midIndex := int(uint(i+j)) >> 1
		if arr[midIndex] > x {
			i = midIndex
		} else {
			j = midIndex
		}
	}
	if Module(arr[i]-x) <= Module(arr[j]-x) {
		return i
	} else {
		return j
	}
}

func Module(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
