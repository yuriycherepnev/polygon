// имеется функция, которая работает неопределенно долго (до 100 секунд)
// написать обертку для этой функции, которая будет прерывать выполнение, если
// функция работает дольше 3 секунд, и возвращать ошибку

package main

import (
	"math/rand"
	"time"
)

func randomTimeWork() {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Second)
}

func predictableTimeWork() {
	ch := make(chan struct{})

	go func() {
		randomTimeWork()
		close(ch)
	}()

	select {
	case <-ch:
	case <-time.After(3 * time.Second):
	}
}

func main() {
	predictableTimeWork()
}
