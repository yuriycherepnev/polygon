package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func randomTimeWorkCh() {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Second)
}

func predictableTimeWorkCh() error {
	ch := make(chan struct{})

	go func() {
		randomTimeWorkCh()
		close(ch)
	}()

	select {
	case <-ch:
		return nil
	case <-time.After(3 * time.Second):
		return errors.New("time is expired")
	}
}

func main() {
	err := predictableTimeWorkCh()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("goroutine worked correctly")
	}
}
