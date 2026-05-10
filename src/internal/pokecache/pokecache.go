package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	cmap map[string]CacheEntry
	m sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		make(map[string]CacheEntry),
		sync.Mutex{},
	}

	go c.ReapLoop(interval)

	return &c
}

func (c *Cache) Add(key string, val []byte) {
	c.m.Lock()
	defer c.m.Unlock()
	c.cmap[key] = CacheEntry{time.Now(),val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.m.Lock()
	defer c.m.Unlock()
	if item, ok := c.cmap[key]; ok {
		return item.val, true
	}
	return nil, false
}

func (c *Cache) ReapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.m.Lock()
		for key, value := range c.cmap {
			if time.Since(value.createdAt) > interval {
				delete(c.cmap, key)
			}
		}
		c.m.Unlock()
	}
}
