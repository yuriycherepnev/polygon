// нужно вывести сколько секунд работала главная горутина mainSeconds
// и сколько суммарно секунд работала каждая горутина отдельно totalWorkSeconds

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// решение на каналах
func randomWaitCh(wg *sync.WaitGroup, workChan chan int) {
	defer wg.Done()
	workSeconds := rand.Intn(9)
	time.Sleep(time.Duration(workSeconds) * time.Second)
	workChan <- workSeconds
}

func main() {
	mainSeconds := 0
	totalWorkSeconds := 0
	start := time.Now()
	workChan := make(chan int)
	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go randomWaitCh(wg, workChan)
	}

	go func() {
		wg.Wait()
		close(workChan)
	}()

	for result := range workChan {
		totalWorkSeconds += result
	}
	elapsed := time.Since(start)
	mainSeconds = int(elapsed.Seconds())

	fmt.Println(mainSeconds)
	fmt.Println(totalWorkSeconds)
}
