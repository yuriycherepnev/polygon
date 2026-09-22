package main

import (
	"fmt"
	"sync"
)

// решение в помощью channels - самое быстрое
type FooBar struct {
	n   int
	sem chan struct{}
}

func NewFooBar(n int) *FooBar {
	f := &FooBar{
		n:   n,
		sem: make(chan struct{}),
	}
	return f
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		printFoo()
		fb.sem <- struct{}{}
		<-fb.sem
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.sem
		printBar()
		fb.sem <- struct{}{}
	}
}

func main() {
	foo := NewFooBar(5)
	var wg sync.WaitGroup

	printFoo := func() { fmt.Print("foo") }
	printBar := func() { fmt.Print("bar") }

	wg.Add(2)

	go func() {
		defer wg.Done()
		foo.Foo(printFoo)
	}()

	go func() {
		defer wg.Done()
		foo.Bar(printBar)
	}()

	wg.Wait()

	fmt.Println()
}
