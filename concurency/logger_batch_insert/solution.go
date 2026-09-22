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
	msgCh  = make(chan string, 100)
	batch  = make([]string, 0, 10)
	wg     = sync.WaitGroup{}
	ticker = time.NewTicker(2 * time.Second)
)

func main() {
	wg.Add(1)
	go logWorker()

	for i := 0; i < 300; i++ {
		logMessage(strconv.Itoa(i))
	}

	wg.Wait()
}

func logMessage(message string) {
	msgCh <- message
}

func insertBatch(messages []string) {
	fmt.Println(messages)
}

func logWorker() {
	defer wg.Done()
	defer ticker.Stop()
	for {
		select {
		case msg := <-msgCh:
			batch = append(batch, msg)
			if len(batch) == 10 {
				insertBatch(batch)
				batch = batch[:0]
			}
		case <-ticker.C:
			if len(batch) > 0 {
				insertBatch(batch)
				batch = batch[:0]
			}
		}
	}
}
