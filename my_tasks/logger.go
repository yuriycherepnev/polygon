package main

import (
	"fmt"
)

type Logger interface {
	Log(message string)
}

type FileLogger struct {
	buffer []string
}

// работать не будет, т.к. чтобы функция с получателем меняла свойства
// структура должна быть передана по ссылке
func (f FileLogger) Log(message string) {
	f.buffer = append(f.buffer, message)
	fmt.Println("saved:", message)
}

// необходимо задействовать интерфейс Logger
type Service struct {
	logger FileLogger
}

func (s Service) Process() {
	s.logger.Log("start process")
}

func main() {
	logger := FileLogger{}

	service := Service{
		logger: logger,
	}

	service.Process()

	fmt.Println(logger.buffer)
}
