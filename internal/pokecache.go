package internal

import (
	"time"
)

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		dataMap: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)
	return c
}
