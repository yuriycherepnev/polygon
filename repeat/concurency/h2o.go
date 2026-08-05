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
	oxygenLock chan struct{}
	hydrogen   chan struct{}
	oxygen     chan struct{}
}

func NewH2O() *H2O {
	h := &H2O{
		oxygenLock: make(chan struct{}, 1),
		hydrogen:   make(chan struct{}, 2),
		oxygen:     make(chan struct{}, 2),
	}
	h.oxygenLock <- struct{}{}
	h.oxygen <- struct{}{}
	h.oxygen <- struct{}{}

	return h
}

func (h *H2O) Hydrogen(releaseHydrogen func()) {
	<-h.hydrogen
	releaseHydrogen()
	h.oxygen <- struct{}{}
}

func (h *H2O) Oxygen(releaseOxygen func()) {
	<-h.oxygenLock
	<-h.oxygen
	<-h.oxygen
	releaseOxygen()
	h.hydrogen <- struct{}{}
	h.hydrogen <- struct{}{}
	h.oxygenLock <- struct{}{}
}
