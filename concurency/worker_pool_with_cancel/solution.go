/*
условие:
операция должна выполняться не более 5 секунд
необходимо:
реализовать функцию processParallel
прокинуть контекст
*/
package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var errorCtxTimeout = errors.New("context timeout")

func processData(ctx context.Context, val int) (int, error) {
	ch := make(chan int)
	go func() {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		close(ch)
	}()
	select {
	case <-ch:
		return val * 2, nil
	case <-ctx.Done():
		return 0, errorCtxTimeout
	}
}

func main() {
	in := make(chan int)
	out := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func() {
		defer close(in)
		for i := range 100 {
			//обработка висящих горутин
			select {
			case in <- i:
			case <-ctx.Done():
				return
			}
		}
	}()

	now := time.Now()
	processParallel(in, out, 5, ctx)

	for val := range out {
		fmt.Println(val)
	}

	fmt.Println(time.Since(now))
}

func processParallel(in <-chan int, out chan<- int, numWorkers int, ctx context.Context) {
	wg := &sync.WaitGroup{}
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, in, out, wg)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
}

func worker(ctx context.Context, in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case number, ok := <-in:
			if !ok {
				return
			}
			result, err := processData(ctx, number)
			if err != nil {
				return
			}
			out <- result
		}
	}
}

/*
Задача решеается с помощью 2 паттернов:
worker pool и cancelation
также паттерн обертка для processData
Алгоритм:
1. Нам необходимо написать воркеров и запустить их
2. Создать контекст и прокинуть его в воркеров для отмены
3. Прокинуть контекст в processData
*/
