package main

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8}
	reverseArray(array)
}

func reverseArray(nums []int) {
	start := 0
	end := len(nums) - 1

	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
