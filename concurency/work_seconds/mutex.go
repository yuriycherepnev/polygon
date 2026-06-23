// нужно вывести сколько секунд работала главная горутина mainSeconds
// и сколько суммарно секунд работала каждая горутина отдельно totalWorkSeconds

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// решение через wg + mutex
func randomWait() int {
	workSeconds := rand.Intn(9)
	time.Sleep(time.Duration(workSeconds) * time.Second)
	return workSeconds
}

func main() {
	mainSeconds := 0
	totalWorkSeconds := 0
	start := time.Now()
	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			workSeconds := randomWait()
			mutex.Lock()
			totalWorkSeconds += workSeconds
			mutex.Unlock()
		}()
	}
	wg.Wait()

	elapsed := time.Since(start)
	mainSeconds = int(elapsed.Seconds())
	fmt.Println(mainSeconds)
	fmt.Println(totalWorkSeconds)
}
