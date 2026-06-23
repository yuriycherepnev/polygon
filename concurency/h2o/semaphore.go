package main

import (
	"fmt"
	"sync"
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
	oLockCh   chan struct{}
	hTokensCh chan struct{}
	oTokensCh chan struct{}
}

func NewH2O() *H2O {
	h := &H2O{
		oLockCh:   make(chan struct{}, 1),
		hTokensCh: make(chan struct{}, 2),
		oTokensCh: make(chan struct{}, 2),
	}
	h.oTokensCh <- struct{}{}
	h.oTokensCh <- struct{}{}
	h.oLockCh <- struct{}{}
	return h
}

func (h *H2O) Hydrogen(releaseHydrogen func()) {
	<-h.hTokensCh
	releaseHydrogen()
	h.oTokensCh <- struct{}{}
}

func (h *H2O) Oxygen(releaseOxygen func()) {
	<-h.oLockCh
	<-h.oTokensCh
	<-h.oTokensCh
	releaseOxygen()
	h.hTokensCh <- struct{}{}
	h.hTokensCh <- struct{}{}
	h.oLockCh <- struct{}{}
}
