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

	for _, src := range sources {
		wg.Add(1)
		go func(s Source) {
			defer wg.Done()
			cars, err := s.Fetch(ctx)
			if err != nil {
				return
			}
			select {
			case results <- cars:
			case <-ctx.Done():
				return
			}
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

	if len(all) == 0 {
		return nil, fmt.Errorf("no cars found")
	}
	return all, nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	cars, err := FetchAll(ctx, sources)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(cars)
}

var (
	sources = []Source{
		fetchSource{
			cars: []Car{
				{ID: 1, Model: "BMW X5", Price: 50000},
				{ID: 2, Model: "Audi A6", Price: 40000},
			},
		},
		fetchSource{
			cars: []Car{
				{ID: 3, Model: "Mercedes E-Class", Price: 60000},
				{ID: 4, Model: "Tesla Model 3", Price: 45000},
			},
		},
		fetchSource{
			cars: []Car{
				{ID: 5, Model: "Toyota Camry", Price: 30000},
			},
		},
	}
)

type fetchSource struct {
	cars []Car
	err  error
}

func (s fetchSource) Fetch(ctx context.Context) ([]Car, error) {
	return s.cars, s.err
}
