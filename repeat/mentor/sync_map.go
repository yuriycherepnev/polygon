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
	fmt.Println(GetOrCreateT("hello", "world"))
	fmt.Println(GetT("hello"))
}

var cache = make(map[string]string)

type CacheT struct {
	CacheM sync.Map
	sync.RWMutex
}

func (c *CacheT) GetOrCreateT(key, value string) string {

	c.Lock()
	value = cache[key]
	c.Unlock()

	if value != "" {
		return value
	}

	c.Lock()
	cache[key] = value
	c.Unlock()

	return value
}

func (c *CacheT) GetT(key string) string {

	v := c.CacheT[key] // c.CacheM.Load(key)
	return v
}
