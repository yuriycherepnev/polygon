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

func (f *FileLogger) Log(message string) {
	f.buffer = append(f.buffer, message)
	fmt.Println("saved:", message)
}

type Service struct {
	logger Logger
}

func (s Service) Process() {
	s.logger.Log("start process")
}

func main() {
	logger := &FileLogger{}
	logger.Log("123")

	service := Service{
		logger: logger,
	}

	service.Process()

	fmt.Println(logger)
}
