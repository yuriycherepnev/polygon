package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	water := "OHHOOOHOHHHHHHH"
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
	mu      sync.Mutex
	cond    *sync.Cond
	counter atomic.Int32
}

func NewH2O() *H2O {
	h := &H2O{
		hCh:     make(chan struct{}, 2),
		oCh:     make(chan struct{}, 2),
		counter: atomic.Int32{},
	}
	h.counter.Store(2)
	h.cond = sync.NewCond(&h.mu)
	h.oCh <- struct{}{}
	return h
}

func (h *H2O) Hydrogen(releaseHydrogen func()) {
	<-h.hCh
	newVal := h.counter.Add(-1)
	if newVal == 0 {
		h.mu.Lock()
		h.cond.Signal()
		h.mu.Unlock()
	}
	releaseHydrogen()
}

func (h *H2O) Oxygen(releaseOxygen func()) {
	<-h.oCh
	releaseOxygen()
	h.hCh <- struct{}{}
	h.hCh <- struct{}{}

	h.mu.Lock()
	for h.counter.Load() > 0 {
		h.cond.Wait()
	}
	h.counter.Store(2)
	h.mu.Unlock()

	h.oCh <- struct{}{}
}
