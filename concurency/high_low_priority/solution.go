package main

import (
	"fmt"
)

func main() {
	high := make(chan int, 10)
	low := make(chan int, 10)

	go func() {
		defer close(high)
		for i := range 10 {
			high <- i * 20
		}
	}()
	go func() {
		defer close(low)
		for i := range 10 {
			low <- i
		}
	}()

	for high != nil || low != nil {
		select {
		case v, ok := <-high:
			if !ok {
				high = nil
				continue
			}
			fmt.Println(v)
		default:
			select {
			case v, ok := <-high:
				if !ok {
					high = nil
					continue
				}
				fmt.Println(v)
			case v, ok := <-low:
				if !ok {
					low = nil
					continue
				}
				fmt.Println(v)
			}
		}
	}
}
