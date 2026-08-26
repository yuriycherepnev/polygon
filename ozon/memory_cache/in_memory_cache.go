package main

import "sync"

type MCache interface {
	Set(k, v string)
	Get(k string) (v string, ok bool)
}

type mCache struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewMCache() MCache {
	return &mCache{
		data: make(map[string]string),
	}
}

func (c *mCache) Set(k, v string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[k] = v
}

func (c *mCache) Get(k string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	v, ok := c.data[k]
	return v, ok
}

func main() {
	memCache := NewCache()
	memCache.Set("123", "123")
}
