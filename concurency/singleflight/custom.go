package main

import (
	"fmt"
	"sync"
	"time"
)

type call struct {
	wg  sync.WaitGroup
	val any
	err error
}

type Group struct {
	mu    sync.Mutex
	calls map[string]*call
}

func NewGroup() *Group {
	return &Group{
		calls: make(map[string]*call),
	}
}

func (g *Group) Do(key string, fn func() (any, error)) (any, error) {
	g.mu.Lock()

	if c, ok := g.calls[key]; ok {
		g.mu.Unlock()
		c.wg.Wait()
		return c.val, c.err
	}
	c := &call{}

	c.wg.Add(1)

	g.calls[key] = c

	g.mu.Unlock()

	c.val, c.err = fn()

	g.mu.Lock()
	delete(g.calls, key)
	g.mu.Unlock()

	c.wg.Done()

	return c.val, c.err
}

func main() {
	group := NewGroup()

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			result, err := group.Do("user:123", func() (any, error) {
				fmt.Println(">>> Выполняем тяжёлую операцию")

				time.Sleep(2 * time.Second)

				return "Yuriy", nil
			})

			if err != nil {
				fmt.Println("goroutine", id, "error:", err)
				return
			}

			fmt.Println("goroutine", id, "result:", result)
		}(i)
	}

	wg.Wait()
}
