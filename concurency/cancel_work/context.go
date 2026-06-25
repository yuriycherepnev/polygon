// имеется функция, которая работает неопределенно долго (до 100 секунд)
// написать обертку для этой функции, которая будет прерывать выполнение, если
// функция работает дольше 3 секунд, и возвращать ошибку

package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func randomTimeWork() {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Second)
}

func predictableTimeWork(ctx context.Context) error {
	go func() {
		randomTimeWork()
	}()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Done")
			return ctx.Err()
		default:
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := predictableTimeWork(ctx)
	fmt.Println(err)
}
