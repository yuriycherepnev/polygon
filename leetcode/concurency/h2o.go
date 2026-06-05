package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	water := "OHOOHOHHHHOHHOHHHH"
	h2o := NewH2O()
	var wg sync.WaitGroup
	for _, ch := range water {
		wg.Add(1)
		switch ch {
		case 'H':
			go func() {
				defer wg.Done()
				h2o.Hydrogen(func() { fmt.Print("H") })
			}()
		case 'O':
			go func() {
				defer wg.Done()
				h2o.Oxygen(func() { fmt.Print("O") })
			}()
		}
	}
	wg.Wait()

	fmt.Println()
}

type H2O struct {
	hCh     chan struct{}
	oCh     chan struct{}
	counter atomic.Int32
}

func NewH2O() *H2O {
	h := &H2O{
		hCh: make(chan struct{}, 2),
		oCh: make(chan struct{}, 2),
	}
	h.oCh <- struct{}{}
	return h
}

func (h *H2O) Hydrogen(releaseHydrogen func()) {
	<-h.hCh
	// releaseHydrogen() outputs "H". Do not change or remove this line.
	releaseHydrogen()
	h.oCh <- struct{}{}
}

func (h *H2O) Oxygen(releaseOxygen func()) {
	<-h.oCh
	// releaseOxygen() outputs "H". Do not change or remove this line.
	releaseOxygen()
	h.hCh <- struct{}{}
	h.hCh <- struct{}{}
}
