/*
Сообщение не сохраняется в буфере, потому что метод Log определён с значением-получателем (value receiver):
func (f FileLogger) Log(...). В Go такой метод работает с копией структуры, и все изменения (append) теряются после возврата.

Как исправить
Сделайте получатель метода Log указателем: func (f *FileLogger) Log(...).
Передавайте в Service указатель: logger := &FileLogger{}.
*/

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

/*
В Go, когда метод определён на значении, внутри метода создаётся копия структуры.
Все изменения полей (например, добавление в buffer) происходят с этой копией и
не влияют на исходный объект после выхода из метода.

Value receiver (получатель-значение) — это всегда копия
Когда метод объявлен с получателем-значением (func (f FileLogger) Log(...)), каждый вызов метода получает копию структуры.
Все изменения внутри метода происходят с этой копией и пропадают после выхода из метода.
*/
