package lru_cache

import (
	"container/list"
	"fmt"
)

type Cache struct {
	capacity int
	items    map[string]*list.Element
	list     *list.List
}

type entry struct {
	key   string
	value string
}

func NewCache(capacity int) *Cache {
	return &Cache{
		capacity: capacity,
		items:    make(map[string]*list.Element),
		list:     list.New(),
	}
}

func (c *Cache) Get(key string) (string, bool) {
	element, ok := c.items[key]
	if !ok {
		return "", false
	}

	c.list.MoveToFront(element)
	return element.Value.(string), true
}

func (c *Cache) Set(key, value string) {
	if element, ok := c.items[key]; ok {
		c.list.MoveToFront(element)
		return
	}
	element := c.list.PushFront(entry{
		key:   key,
		value: value,
	})
	c.items[key] = element

	if c.list.Len() > c.capacity {
		oldest := c.list.Back()
		oldestEntry := oldest.Value.(entry)
		delete(c.items, oldestEntry.key)
		c.list.Remove(oldest)
	}
}

func (c *Cache) Delete(key string) {
	element, ok := c.items[key]
	if !ok {
		return
	}
	delete(c.items, key)
	c.list.Remove(element)
}

func main() {
	cache := NewCache(2)
	cache.Set("a", "Apple")
	fmt.Println(cache.Get("a"))
}
