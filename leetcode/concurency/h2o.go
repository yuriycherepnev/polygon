package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	water := "OOOHHH"
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
}

type H2O struct {
	wgHydrogen sync.WaitGroup
	wgOxygen   sync.WaitGroup
	counter    atomic.Int32
}

func NewH2O() *H2O {
	h := &H2O{
		wgHydrogen: sync.WaitGroup{},
		wgOxygen:   sync.WaitGroup{},
		counter:    atomic.Int32{},
	}
	h.wgHydrogen.Add(1)
	return h
}

func (h *H2O) Hydrogen(releaseHydrogen func()) {
	h.wgHydrogen.Wait()
	// releaseHydrogen() outputs "H". Do not change or remove this line.
	releaseHydrogen()
	h.wgOxygen.Done()
	h.wgHydrogen.Add(1)
}

func (h *H2O) Oxygen(releaseOxygen func()) {
	h.wgOxygen.Wait()
	// releaseOxygen() outputs "H". Do not change or remove this line.
	releaseOxygen()
	h.wgOxygen.Add(1)
	h.wgHydrogen.Done()
}
