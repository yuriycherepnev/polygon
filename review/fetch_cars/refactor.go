package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Car struct {
	ID    int
	Model string
	Price int
}

type Source interface {
	Fetch(ctx context.Context) ([]Car, error)
}

func FetchAll(ctx context.Context, sources []Source) ([]Car, error) {
	var wg sync.WaitGroup
	results := make(chan []Car)
	var total int

	for _, src := range sources {
		wg.Add(1)
		go func(s Source) {
			defer wg.Done()
			cars, err := s.Fetch(ctx)
			if err != nil {
				return
			}
			total++
			results <- cars
		}(src)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var all []Car
	for batch := range results {
		all = append(all, batch...)
	}

	if total == 0 {
		return nil, fmt.Errorf("no cars found")
	}
	return all, nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	sources := []Source{}
	cars, err := FetchAll(ctx, sources)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("cars:", cars)
}

var ()

type FetchSource struct {
}

func (f FetchSource) Fetch(ctx context.Context) ([]Car, error) {
	return nil, nil
}
