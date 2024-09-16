package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mu    *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(duration time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mu:    &sync.Mutex{},
	}

	go c.reapLoop(duration)

	return c
}

// Debug function
// Get amount of all items in cache
// func (c *Cache) Len() int {
// 	c.mu.Lock()
//     defer c.mu.Unlock()
//
//     return len(c.cache)
// }

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()

	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

    c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
    defer c.mu.Unlock()

	val, ok := c.cache[key]

	return val.val, ok
}

// Debug function
// Print all locations in the map
func (c *Cache) GetAll() {
	c.mu.Lock()

    fmt.Println("DEBUG locations")
    for k, v := range c.cache {
        fmt.Printf("%v %v\n", k, v)
    }
    fmt.Println("DEBUG end locations")

    c.mu.Unlock()
}


func (c *Cache) reapLoop(duration time.Duration) {
	ticker := time.NewTicker(duration)
	for {
		select {
		case <-ticker.C:
            c.removeExpired(time.Now(), duration)
		}
	}
}

func (c *Cache) removeExpired(now time.Time, interval time.Duration) {
	c.mu.Lock()

	for key, val := range c.cache {
		if val.createdAt.Before(now.Add(-interval)) {
			delete(c.cache, key)
		}
	}

    c.mu.Unlock()
}
