package proteincache

import (
	"fmt"
	"sync"
	"time"
)

// Cache -
type Cache struct {
	cache map[string]cacheEntry
	queries map[string]string
	mux   *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// NewCache -
func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		queries: make(map[string]string),
		mux:   &sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}

// Add -
func (c *Cache) Add(key string, value []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}
}
// AddRid -
func (c *Cache) AddRid(key string, value string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.queries[key] = value
}

// PrintRids -
func (c *Cache) PrintRids() {
	for k, v := range c.queries {
		fmt.Printf("Query: %s -> RID: %s\n ", k, v)
	}
}

// Get -
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	val, ok := c.cache[key]
	return val.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k, v := range c.cache {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.cache, k)
		}
	}
}
