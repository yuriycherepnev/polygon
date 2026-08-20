package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	urls := []string{
		"https://www.ozon.ru/",
		"https://music.yandex.ru/",
		"https://www.ozon.ru/",
		"https://music.yandex.ru/",
		"https://www.ozon.ru/",
		"https://music.yandex.ru/",
	}
	jobs := make(chan string, len(urls))
	results := make(chan string, len(urls))
	wg := &sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(ctx, jobs, results, wg)
	}

	go func() {
		for i, url := range urls {
			if i == 2 {
				cancel()
				break
			}
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
