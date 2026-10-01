package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func unpredictableFunc() int64 {
	rnd := rand.Int63n(5000)
	time.Sleep(time.Duration(rnd) * time.Millisecond)
	return rnd
}

func predictableFunc(ctx context.Context) (int64, error) {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()
	result := make(chan int64, 1)

	go func() {
		defer close(result)
		result <- unpredictableFunc()
	}()

	select {
	case <-ctx.Done():
		return 0, nil
	case v, _ := <-result:
		return v, nil
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	fmt.Println("started")
	fmt.Println(predictableFunc(ctx))
}
