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
}

type Foo struct {
	first  chan struct{}
	second chan struct{}
}

func NewFoo() *Foo {
	return &Foo{
		first:  make(chan struct{}),
		second: make(chan struct{}),
	}
}

func (f *Foo) First(printFirst func()) {
	printFirst()
	f.first <- struct{}{}
}

func (f *Foo) Second(printSecond func()) {
	<-f.first
	printSecond()
	f.second <- struct{}{}
}

func (f *Foo) Third(printThird func()) {
	<-f.second
	printThird()
}
