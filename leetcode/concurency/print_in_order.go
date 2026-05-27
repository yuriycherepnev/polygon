/*
sync.WaitGroup — это просто структура из 3 слов (счётчик и семафор).
Операции Add, Done, Wait используют атомарные инструкции и легковесное ожидание (futex).
Нет дополнительных аллокаций, нет копирования данных.
-----
Канал — это сложная структура (hchan) с:
внутренним мьютексом,
кольцевой очередью,
двумя очередями ожидания (для отправки и получения),
может требовать аллокации при создании и иногда при операциях.
В вашем коде канал chan int — каждый раз при отправке/получении значения типа int (8 байт)
происходит его копирование в очередь канала.
-----
Память
WaitGroup: 0 дополнительных аллокаций (кроме самой структуры в Foo).

Каналы: две выделенные кучи hchan (с блокировками, очередями, внутренними буферами).
Плюс возможные временные аллокации, если канал использует буфер
(но у вас небуферизированный → всё равно есть структуры ожидания).
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	foo := NewFoo()
	var wg sync.WaitGroup

	printFirst := func() { fmt.Print("first") }
	printSecond := func() { fmt.Print("second") }
	printThird := func() { fmt.Print("third") }

	wg.Add(3)

	go func() {
		defer wg.Done()
		foo.Third(printThird)
	}()

	go func() {
		defer wg.Done()
		foo.Second(printSecond)
	}()

	go func() {
		defer wg.Done()
		foo.First(printFirst)
	}()

	wg.Wait()
	fmt.Println()
}

// решение через wait group
type Foo struct {
	first  sync.WaitGroup
	second sync.WaitGroup
}

func NewFoo() *Foo {
	f := &Foo{
		first:  sync.WaitGroup{},
		second: sync.WaitGroup{},
	}
	// Устанавливаем счётчики в 1
	f.first.Add(1)
	f.second.Add(1)
	return f
}

func (f *Foo) First(printFirst func()) {
	// Do not change this line
	printFirst()
	f.first.Done()
}

func (f *Foo) Second(printSecond func()) {
	f.first.Wait()
	/// Do not change this line
	printSecond()
	f.second.Done()
}

func (f *Foo) Third(printThird func()) {
	f.second.Wait()
	// Do not change this line
	printThird()
}

// решение через каналы
/*
type Foo struct {
	firstChan  chan int
	secondChan chan int
}

func NewFoo() *Foo {
	return &Foo{
		firstChan:  make(chan int),
		secondChan: make(chan int),
	}
}

func (f *Foo) First(printFirst func()) {
	// Do not change this line
	printFirst()
	f.firstChan <- 1
}

func (f *Foo) Second(printSecond func()) {
	number := <-f.firstChan
	/// Do not change this line
	printSecond()
	f.secondChan <- number
}

func (f *Foo) Third(printThird func()) {
	<-f.secondChan
	// Do not change this line
	printThird()
}
*/
