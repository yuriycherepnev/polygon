/*
есть логгер сообщений, куда приходит лог с разных сервисов.
Нужно не меняя сигнатуры функций
сделать сохранение сообщений пачками по 10 шутк, либо каждые 2 секунды
*/

package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var (
	messages = make(chan string, 100)
	wg       = sync.WaitGroup{}
	ticker   = time.NewTicker(2 * time.Second)
)

func main() {
	wg.Add(1)
	go logWorker()

	for i := 0; i < 300; i++ {
		logMessage(strconv.Itoa(i))

		//if (i % 50) == 0 {
		//	time.Sleep(2 * time.Second)
		//}
	}

	wg.Wait()
}

func logMessage(message string) {
	messages <- message
}

func insertBatch(messages []string) {
	fmt.Println(messages)
}

func logWorker() {
	defer wg.Done()
	defer ticker.Stop()

	batch := make([]string, 0, 10)

	for {
		select {
		case message := <-messages:
			batch = append(batch, message)

			if len(batch) == 10 {
				fmt.Println("batch")
				insertBatch(batch)
				batch = batch[:0]
			}

		case <-ticker.C:
			if len(batch) > 0 {
				fmt.Println("ticker")
				insertBatch(batch)
				batch = batch[:0]
			}
		}
	}
}
