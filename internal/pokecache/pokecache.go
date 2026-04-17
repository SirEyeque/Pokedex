package pokecache

import (
	"time"
	"sync"
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
	return &c
}

func Add(c *Cache, key string, val []byte) {
	c.cmap[key] = CacheEntry{time.Now(),val}
}

func Get(c *Cache, key string) ([]byte, bool) {
	if item, ok := c.cmap[key]; ok {
		return item.val, true
	}
	return nil, false
}

func reapLoop() {
}
