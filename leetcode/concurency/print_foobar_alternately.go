package main

import (
	"fmt"
	"sync"
)

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

func main() {
	foo := NewFooBar(1)
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
