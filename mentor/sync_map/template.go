/*
Разработчик дал на ревью код своего нового кэша, нам необходимо провести код-ревью
 Кэш будет использоваться под высокой нагрузкой в проде
 Частота записи/чтения 20%/80% соответственно
*/

package main

/*
import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println(GetOrCreateT("hello", "world"))
	fmt.Println(GetT("hello"))
}

// здесь cache это глобальная переменная уровня пакета.
// Она инициализируется один раз.
var cache = make(map[string]string)

type CacheT struct {
	CacheM sync.Map
	sync.RWMutex
}

// GetOrCreateT checks whether the key exists.
// If it does not exist, it creates a new value.
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
*/

/*
в изначальном шаблоне есть 1 глобальная проблема:
попытка 2 разных подхода в 1.
Т.е. нужно использовать либо обычную мапу + rwMutex
либо syncMap которая внутри себя уже реализует атомики
для записи и чтения.

problems:

GetT
1. в syncMap такое обращение не сработает:
v := c.CacheT[key]

GetOrCreateT
здесь используется обычная мапа + рвмьютекс

1. value - затрется пустой строкой в GetOrCreateT

2. при проверке существования ключа в map
надо использовать value, ok := cache[key]

3. Между проверкой и записью не должно быть пространства
c.Lock()
c.Unlock()
это должна быть одна логическая операция

вопросы:
1. насколько хорошим тоном считается анонимное встраивание типа?

type Cache struct {
	cacheMap sync.Map
	sync.RWMutex
}
*/
