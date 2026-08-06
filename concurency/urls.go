package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

type urlResponse struct {
	url string
	len int
	err error
}

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
	results := make(chan urlResponse, len(urls))
	wg := &sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(jobs, results, wg)
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
		fmt.Println(message.url, message.len, message.err)
	}
}

func worker(jobs chan string, results chan urlResponse, wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range jobs {
		resp, err := http.Get(url)
		if err != nil {
			results <- urlResponse{
				url: url,
				err: err,
			}
			continue
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			results <- urlResponse{
				url: url,
				err: err,
			}
			continue
		}

		resp.Body.Close()

		results <- urlResponse{
			url: url,
			len: len(body),
		}
	}
}
