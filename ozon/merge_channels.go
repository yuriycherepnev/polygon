package main

import "sync"

func main() {
	chnCount := 30
	channels := make([]<-chan int, chnCount)

	for i := 0; i < chnCount; i++ {
		channels[i] = make(chan int)
	}

	out := merge(channels...)
	_ = out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	results := make(chan int)

	for _, channel := range cs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range channel {
				results <- v
			}
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return results
}
