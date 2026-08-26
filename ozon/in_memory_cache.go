package main

import "sync"

type Cache interface {
	Set(k, v string)
	Get(k string) (v string, ok bool)
}

type cache struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewMCache() Cache {
	return &cache{
		data: make(map[string]string),
	}
}

func (c *cache) Set(k, v string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[k] = v
}

func (c *cache) Get(k string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	v, ok := c.data[k]
	return v, ok
}

func main() {
	mCache := NewMCache()
	mCache.Set("123", "123")
}
