package main

import (
	"fmt"
	"sync"
)

// решение в помощью channels - самое быстрое
type FooBar struct {
	n   int
	chn chan struct{}
}

func NewFooBar(n int) *FooBar {
	f := &FooBar{
		n:   n,
		chn: make(chan struct{}),
	}
	return f
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		// printFoo() outputs "foo". Do not change or remove this line.
		printFoo()
		fb.chn <- struct{}{}
		<-fb.chn
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.chn
		// printBar() outputs "bar". Do not change or remove this line.
		printBar()
		fb.chn <- struct{}{}
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

// FooBar решение через атомики
/*
Функция runtime.Gosched() в языке Go принудительно возвращает текущую горутину
в глобальную очередь и освобождает процессор.
*/
/*
type FooBar struct {
	n     int
	isBar atomic.Bool
}

func NewFooBar(n int) *FooBar {
	return &FooBar{
		n:     n,
		isBar: atomic.Bool{},
	}
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		for fb.isBar.Load() {
			runtime.Gosched()
		}
		// printFoo() outputs "foo". Do not change or remove this line.
		printFoo()
		fb.isBar.Store(true)
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		for !fb.isBar.Load() {
			runtime.Gosched()
		}
		// printBar() outputs "bar". Do not change or remove this line.
		printBar()
		fb.isBar.Store(false)
	}
}
*/

// решение в помощью wg
/*
type FooBar struct {
	n   int
	foo sync.WaitGroup
	bar sync.WaitGroup
}

func NewFooBar(n int) *FooBar {
	f := &FooBar{n: n}
	f.foo.Add(1)
	return f
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		fb.bar.Wait()
		// printFoo() outputs "foo". Do not change or remove this line.
		printFoo()
		fb.bar.Add(1)
		fb.foo.Done()
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		fb.foo.Wait()
		// printBar() outputs "bar". Do not change or remove this line.
		printBar()
		fb.bar.Done()
		if i+1 < fb.n {
			fb.foo.Add(1)
		}
	}
}
*/
