package main

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	numbers = append(numbers[:1], numbers[2:]...)
}
