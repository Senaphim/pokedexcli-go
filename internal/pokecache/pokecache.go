package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	value     []byte
}

type Cache struct {
	Entries map[string]cacheEntry
	mu      *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		Entries: make(map[string]cacheEntry),
		mu:      &sync.Mutex{},
	}
	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Entries[key] = cacheEntry{
		createdAt: time.Now().Local(),
		value:     val,
	}

	return
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	var ret []byte
	entry, present := c.Entries[key]
	if present {
		ret = entry.value
	}
	return ret, present
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// Use range on ticker to constantly pull values from the queue
	for range ticker.C {
		// Could probably use time sent by ticker here, but using .Now() is
		// probaby safer. Converting to local as this matches createdat
		now := time.Now().Local()
		c.mu.Lock()

		// Loop over entries and check if they are too old. Using methods to do
		// comparison as time is wierd
		for i, entry := range c.Entries {
			if entry.createdAt.Before(now.Add(-interval)) {
				delete(c.Entries, i)
			}
		}

		// Cant use defer statement as will not reach return to call defer
		c.mu.Unlock()
	}
	return
}
