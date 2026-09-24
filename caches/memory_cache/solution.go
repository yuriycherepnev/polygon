package main

import (
	"fmt"
	"sync"
)

type Cache interface {
	Set(k, v string)
	Get(k string) (v string, ok bool)
}

type cache struct {
	mu sync.RWMutex
	m  map[string]string
}

func (c *cache) Set(k, v string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[k] = v
}

func (c *cache) Get(k string) (v string, ok bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok = c.m[k]
	return v, ok
}

func NewCache() Cache {
	c := &cache{
		m: make(map[string]string),
	}
	return c
}

func main() {
	mCache := NewCache()
	mCache.Set("foo", "bar")
	fmt.Println(mCache.Get("foo"))
}
