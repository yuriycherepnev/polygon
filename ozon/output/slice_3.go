package main

import "time"

func main() {
	timeStart := time.Now()
	_, _ = <-worker(), <-worker()
	println(int(time.Since(timeStart).Seconds()))
}

func worker() chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- 1
	}()
	return ch
}

// вывод 6
// чтение из канала блокирующая операция
