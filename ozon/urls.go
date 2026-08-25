package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	urls := []string{
		"https://example.com",
		"https://golang.org",
		"https://google.com",
		"https://example.com",
	}
	jobs := make(chan string, len(urls))
	results := make(chan string, len(urls))
	wg := &sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(ctx, jobs, results, wg)
	}

	go func() {
		for _, url := range urls {
			jobs <- url
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}
}

func worker(ctx context.Context, jobs chan string, result chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case url, ok := <-jobs:
			if !ok {
				return
			}
			_, err := http.Get(url)
			if err != nil {
				result <- url + " not ok"
			} else {
				result <- url + " ok"
			}
		}
	}
}
