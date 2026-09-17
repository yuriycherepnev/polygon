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

		go func(id int) {
			defer wg.Done()

			result, err, shared := group.Do("user:123", func() (any, error) {
				fmt.Println(">>> Выполняем запрос в БД")

				// Имитируем медленный запрос
				time.Sleep(2 * time.Second)

				return "Yuriy", nil
			})

			if err != nil {
				fmt.Println("goroutine", id, "error:", err)
				return
			}

			fmt.Printf(
				"goroutine %d: result=%v, shared=%v\n",
				id,
				result,
				shared,
			)
		}(i)
	}

	wg.Wait()
}
