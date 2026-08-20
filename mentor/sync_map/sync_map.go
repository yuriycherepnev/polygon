// Разработчик дал на ревью код своего нового кэша, нам необходимо провести код-ревью
// Кэш будет использоваться под высокой нагрузкой в проде
// Частота записи/чтения 20%/80% соответственно

package main

import (
	"fmt"
	"sync"
)

func main() {
	cacheMap := Cache[string, string]{}
	fmt.Println(cacheMap.GetOrCreate("hello", "world"))
	fmt.Println(cacheMap.Get("hello"))
}

type Cache[K comparable, V any] struct {
	cacheMap sync.Map
}

func (c *Cache[K, V]) GetOrCreate(key, value string) (V, bool) {
	store, ok := c.cacheMap.LoadOrStore(key, value)
	return store.(V), ok
}

func (c *Cache[K, V]) Get(key string) (V, bool) {
	v, ok := c.cacheMap.Load(key)
	if ok {
		return v.(V), true
	}
	var zero V
	return zero, false
}
