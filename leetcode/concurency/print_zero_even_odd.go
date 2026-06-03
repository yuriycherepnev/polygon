package main

import (
	"fmt"
	"sync"
)

func main() {
	evenOdd := NewZeroEvenOdd(3)
	wg := sync.WaitGroup{}

	printZero := func(n int) { fmt.Print(n) }
	printEven := func(n int) { fmt.Print(n) }
	printOdd := func(n int) { fmt.Print(n) }

	wg.Add(3)

	go func() {
		defer wg.Done()
		evenOdd.Zero(printZero)
	}()

	go func() {
		defer wg.Done()
		evenOdd.Even(printEven)
	}()

	go func() {
		defer wg.Done()
		evenOdd.Odd(printOdd)
	}()

	wg.Wait()

	fmt.Println()

}

type ZeroEvenOdd struct {
	n    int
	zero sync.WaitGroup
	even sync.WaitGroup
	odd  sync.WaitGroup
}

func NewZeroEvenOdd(n int) *ZeroEvenOdd {
	zeo := &ZeroEvenOdd{
		n:    n,
		zero: sync.WaitGroup{},
		odd:  sync.WaitGroup{},
		even: sync.WaitGroup{},
	}
	zeo.zero.Add(1)
	zeo.odd.Add(1)

	return zeo
}

func (z *ZeroEvenOdd) Zero(printNumber func(int)) {
	for i := 1; i <= z.n; i++ {
		z.even.Wait()
		printNumber(0)
		z.zero.Done()
		z.even.Add(1)
	}
}

func (z *ZeroEvenOdd) Odd(printNumber func(int)) {
	for i := 1; i <= z.n; i++ {
		z.zero.Wait()
		if i%2 != 0 {
			printNumber(i)
		}
		z.odd.Done()
		z.zero.Add(1)
	}
}

func (z *ZeroEvenOdd) Even(printNumber func(int)) {
	for i := 1; i <= z.n; i++ {
		z.odd.Wait()
		if i%2 == 0 {
			printNumber(i)
		}
		z.even.Done()
		z.odd.Add(1)
	}
}
