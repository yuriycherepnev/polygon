package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	money := atomic.Int32{}
	count := atomic.Int32{}
	groups := 1000
	mutex := sync.Mutex{}

	go func() {
		for i := 0; i < groups; i++ {
			mutex.Lock()
			m := money.Load()
			c := count.Load()
			mutex.Unlock()

			if m != c {
				fmt.Println(m, c)
			}
		}
	}()

	wg := sync.WaitGroup{}

	wg.Add(groups)
	for range groups {
		go func() {

			defer wg.Done()

			mutex.Lock()
			money.Add(1)
			count.Add(1)
			mutex.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println(money.Load())
}
