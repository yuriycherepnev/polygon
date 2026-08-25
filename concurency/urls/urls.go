package main

import (
	"fmt"
	"net/http"
	"sync"
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

	for i := 0; i < 3; i++ {
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
		fmt.Println(message)
	}
}

func worker(jobs chan string, result chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for url := range jobs {
		resp, err := http.Get(url)

		if err != nil {
			result <- url + " not ok"
			continue
		}

		resp.Body.Close()
		result <- url + " ok"
	}
}
