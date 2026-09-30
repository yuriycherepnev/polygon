package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

func main() {
	var group singleflight.Group
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			result, _, _ := group.Do("user:123", func() (any, error) {
				num := heavyCalc()
				return num, nil
			})
			fmt.Println(result.(int))
		}()
	}

	wg.Wait()
}

func heavyCalc() int {
	fmt.Println("heavyCalc")
	time.Sleep(2 * time.Second)
	return 123
}
