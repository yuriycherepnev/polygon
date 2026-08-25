package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	urls := []string{
		"https://example.com",
		"https://golang.org",
		"https://google.com",
		"https://example.com",
		"https://golang.org",
		"https://google.com",
		"https://example.com",
		"https://golang.org",
		"https://google.com",
	}

	jobs := make(chan string, len(urls))
	results := make(chan string, len(urls))
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go ctxWorker(ctx, jobs, results, wg)
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

	for message := range results {
		fmt.Println(message)
	}
}

func ctxWorker(ctx context.Context, jobs chan string, result chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	client := &http.Client{}

	for {
		select {
		case <-ctx.Done():
			return
		case url, ok := <-jobs:
			if !ok {
				return
			}
			//resp, err := http.Get(url)
			request, err := http.NewRequestWithContext(
				ctx,
				http.MethodGet,
				url,
				nil,
			)
			resp, err := client.Do(request)

			if err != nil {
				result <- url + " not ok"
				continue
			}

			resp.Body.Close()
			result <- url + " ok"
		}
	}

}
