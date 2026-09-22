/*
Написать конструкцию `select` так, чтобы сообщения из `high` обрабатывались в первую очередь,
сообщения из `low` обрабатывались если high пуст
*/

package main

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
}
