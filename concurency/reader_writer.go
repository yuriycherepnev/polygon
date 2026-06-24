/*
Написать 3 функции:
writer - генерит числа от 1 до 10
doubler - умножает числа на 2, имитируя работу (500ms)
reader - читает и выводит на экран
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	writeCh := writer()
	doubleCh := doubler(writeCh)
	reader(doubleCh)
}

func writer() <-chan int {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	return ch
}

func doubler(ch <-chan int) <-chan int {
	doubleCh := make(chan int)

	go func() {
		for number := range ch {
			time.Sleep(1 * time.Second / 2)
			doubleCh <- number * 2
		}
		close(doubleCh)
	}()

	return doubleCh
}

func reader(ch <-chan int) {
	for message := range ch {
		fmt.Println(message)
	}
}
