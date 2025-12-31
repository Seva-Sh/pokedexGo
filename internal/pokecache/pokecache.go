package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	CacheMap map[string]CacheEntry
	Mu       *sync.Mutex
	Interval time.Duration
}

type CacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		CacheMap: make(map[string]CacheEntry),
		Mu:       &sync.Mutex{},
		Interval: interval,
	}

	go c.ReapLoop()

	return c
}

func (c *Cache) Add(key string, val []byte) {
	newCacheEntry := CacheEntry{
		CreatedAt: time.Now(),
		Val:       val,
	}
	c.Mu.Lock()
	c.CacheMap[key] = newCacheEntry
	c.Mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mu.Lock()
	defer c.Mu.Unlock()

	entry, exists := c.CacheMap[key]
	if !exists {
		return []byte{}, false
	}
	return entry.Val, true
}

func (c *Cache) ReapLoop() {
	ticker := time.NewTicker(c.Interval)

	for range ticker.C {
		c.Mu.Lock()
		for key, val := range c.CacheMap {
			timeSince := time.Since(val.CreatedAt)
			if timeSince > c.Interval {
				delete(c.CacheMap, key)
			}
		}
		c.Mu.Unlock()
	}
}
