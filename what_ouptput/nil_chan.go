package main

import "fmt"

func main() {
	var ch chan int

	go func() {
		for i := range 5 {
			ch <- i
		}
	}()

	for val := range ch {
		fmt.Println(val)
	}
}

/*
будет паника, канал не инициализирован
Чтобы исправить надо добавить make и close
ch := make(chan int)
close(ch)
*/
