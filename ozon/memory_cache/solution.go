package main

import "sync"

const shardCount = 16

type Cache interface {
	Set(k, v string)
	Get(k string) (v string, ok bool)
}

type cache struct {
	shards [shardCount]shard
}

type shard struct {
	mu sync.RWMutex
	m  map[string]string
}

func (c *cache) Set(k, v string) {
	s := c.GetShard(1)

	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[k] = v
}

func (c *cache) Get(k string) (v string, ok bool) {
	s := c.GetShard(1)

	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok = s.m[k]
	return v, ok
}

func (c *cache) GetShard(i int) *shard {
	return &c.shards[i]
}

func NewCache() Cache {
	c := &cache{}

	for i := 0; i < shardCount; i++ {
		c.shards[i].m = make(map[string]string)
	}

	return c
}

func main() {

}
