// операция должна выполняться не более 5 секунд

package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func processData(val int) int {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	return val * 2
}

func main() {
	in := make(chan int)
	out := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func() {
		for i := range 100 {
			in <- i
		}
		close(in)
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
		go worker(in, out, wg, ctx)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
}

func worker(in <-chan int, out chan<- int, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case number, ok := <-in:
			if !ok {
				return
			}
			select {
			case <-ctx.Done():
				return
			case out <- processData(number):
			}
		}
	}
}

/*
Задача решеается с помощью 2 паттернов:
worker pool и cancelation
Алгоритм:
1. Сначала нам необходимо написать воркеров и запустить их
2. Затем создать контекст и прокинуть его в воркеров для отмены
*/
