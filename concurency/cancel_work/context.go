package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func randomTimeWorkCh() {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Second)
}

func predictableTimeWorkCh(ctx context.Context) error {
	ch := make(chan struct{})

	go func() {
		randomTimeWorkCh()
		close(ch)
	}()

	select {
	case <-ch:
		return nil
	case <-ctx.Done():
		return errors.New("time is expired")
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := predictableTimeWorkCh(ctx)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("goroutine worked correctly")
	}
}
