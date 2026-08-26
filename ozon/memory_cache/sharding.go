package main

import (
	"hash/crc32"
	"sync"
)

const shardsCount = 16

type Cache interface {
	Set(k, v string)
	Get(k string) (v string, ok bool)
}

type shard struct {
	mu   sync.RWMutex
	data map[string]string
}

type cache struct {
	shards [shardsCount]shard
}

func NewCache() Cache {
	c := &cache{}

	for i := range c.shards {
		c.shards[i].data = make(map[string]string)
	}

	return c
}

func (c *cache) Set(k, v string) {
	cShard := c.getShard(k)

	cShard.mu.Lock()
	defer cShard.mu.Unlock()

	cShard.data[k] = v
}

func (c *cache) Get(k string) (string, bool) {
	cShard := c.getShard(k)

	cShard.mu.RLock()
	defer cShard.mu.RUnlock()

	v, ok := cShard.data[k]
	return v, ok
}

func (c *cache) getShard(k string) *shard {
	hash := crc32.ChecksumIEEE([]byte(k))
	index := hash % shardsCount
	return &c.shards[index]
}
