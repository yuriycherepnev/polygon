/*
Разработчик дал на ревью код своего нового кэша, нам необходимо провести код-ревью
 Кэш будет использоваться под высокой нагрузкой в проде
 Частота записи/чтения 20%/80% соответственно
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	cacheMap := CacheT{}

	fmt.Println(cacheMap.GetOrCreateT("hello", "world"))
	fmt.Println(cacheMap.GetT("hello"))
}

var cache = make(map[string]string)

type CacheT struct {
	CacheM sync.Map
	sync.RWMutex
}

func (c *CacheT) GetOrCreateT(key, value string) string {
	c.RLock()
	v, ok := cache[key]
	c.RUnlock()
	if ok {
		return v
	}

	c.Lock()
	cache[key] = value
	c.Unlock()

	return value
}

func (c *CacheT) GetT(key string) string {
	c.RLock()
	defer c.RUnlock()
	v, ok := cache[key]
	if ok {
		return v
	}
	return ""
}
