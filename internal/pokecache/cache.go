package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntry map[string]cacheEntry
	mu         *sync.Mutex
}
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		cacheEntry: map[string]cacheEntry{},
		mu:         &sync.Mutex{},
	}
	go cache.reapLoop(interval)
	return cache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.cacheEntry[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}
func (cache Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	if entry, ok := cache.cacheEntry[key]; ok {
		return entry.val, ok
	}
	return nil, false
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for t := range ticker.C {
		for key, entry := range cache.cacheEntry {
			if t.Sub(entry.createdAt) > interval {
				delete(cache.cacheEntry, key)
			}
		}
	}
}
