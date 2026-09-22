/*
Написать конструкцию `select` так, чтобы сообщения из `high` обрабатывались в первую очередь,
сообщения из `low` обрабатывались если high пуст
*/

package main

import (
	"sync"
)

func main() {
	high := make(chan int, 10)
	low := make(chan int, 10)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		defer close(high)
		for i := range 10 {
			high <- i * 20
		}
	}()
	go func() {
		defer wg.Done()
		defer close(low)
		for i := range 10 {
			low <- i
		}
	}()

	wg.Wait()
}
