package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

type cache struct {
	mu sync.RWMutex
	m  map[string]string
}

func (c *cache) Set(k, v string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[k] = v
}

func (c *cache) Get(k string) (v string, ok bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok = c.m[k]
	return v, ok
}

func WeatherForecast() int {
	time.Sleep(1 * time.Second)
	return rand.Intn(70) - 30
}

func main() {
	var group singleflight.Group
	mCache := cache{
		m: make(map[string]string),
	}

	http.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		temperature, ok := mCache.Get("t")
		if !ok {
			result, _, _ := group.Do("temperature", func() (any, error) {
				t := WeatherForecast()
				return t, nil
			})
			t := result.(int)
			mCache.Set("t", strconv.Itoa(t))
		}
		fmt.Fprintf(w, "{\"temperature\":%d}\n", temperature)
	})

	if err := http.ListenAndServe(":3333", nil); err != nil {
		panic(err)
	}
}
