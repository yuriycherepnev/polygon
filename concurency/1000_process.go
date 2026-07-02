//Написать код который обрабатывает 1000 функций process,
//но не более 3 одновременно, не используя worker pool паттерн

package main

import (
	"fmt"
	"sync"
	"time"
)

func process(counter int) {
	fmt.Println("Start", counter)
	time.Sleep(1 * time.Second)
	fmt.Println("Stop", counter)
}

func main() {
	semCh := make(chan struct{}, 3)
	wg := &sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go sWorker(i, semCh, wg)
	}

	wg.Wait()
}

func sWorker(i int, semCh chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	semCh <- struct{}{}
	process(i)
	<-semCh
}
