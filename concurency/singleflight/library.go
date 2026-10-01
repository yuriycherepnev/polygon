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

	for i := range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			result, err, _ := group.Do("user:123", func() (any, error) {
				num, err := heavyCalc()
				return num, err
			})
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(result, i)
		}()
	}

	wg.Wait()
}

func heavyCalc() (int, error) {
	fmt.Println("heavyCalc")
	time.Sleep(2 * time.Second)
	return 123, nil
}
