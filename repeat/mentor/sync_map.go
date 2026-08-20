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
	cacheM := CacheM{}

	fmt.Println(cacheM.GetOrCreate("hello", "world"))
	fmt.Println(cacheM.Get("hello"))
}

var cache = make(map[string]string)

type CacheM struct {
	mtx  sync.RWMutex
	sMap sync.Map
}

// GetOrCreate проверяет существование ключа key
// Если такого нет, то создает новое значение
func (c *CacheM) GetOrCreate(key, value string) string {
	c.mtx.RLock()
	value, ok := cache[key]
	c.mtx.RUnlock()
	if ok {
		return value
	}
	c.mtx.Lock()
	defer c.mtx.Unlock()
	cache[key] = value
	return value
}

func (c *CacheM) Get(key string) string {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	v := cache[key]
	return v
}
