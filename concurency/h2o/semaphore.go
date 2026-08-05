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
	oxygenLockCh chan struct{}
	hydrogenCh   chan struct{}
	oxygenCh     chan struct{}
}

func NewH2O() *H2O {
	h := &H2O{
		oxygenLockCh: make(chan struct{}, 1),
		hydrogenCh:   make(chan struct{}, 2),
		oxygenCh:     make(chan struct{}, 2),
	}
	h.oxygenLockCh <- struct{}{}
	h.oxygenCh <- struct{}{}
	h.oxygenCh <- struct{}{}
	return h
}

func (h *H2O) Hydrogen(releaseHydrogen func()) {
	<-h.hydrogenCh
	releaseHydrogen()
	h.oxygenCh <- struct{}{}
}

func (h *H2O) Oxygen(releaseOxygen func()) {
	<-h.oxygenLockCh
	<-h.oxygenCh
	<-h.oxygenCh
	releaseOxygen()
	h.hydrogenCh <- struct{}{}
	h.hydrogenCh <- struct{}{}
	h.oxygenLockCh <- struct{}{}
}
