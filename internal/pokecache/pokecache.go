package internal

import (
	"time"
)

func NewCache(minuteInterval time.Duration, secInterval time.Duration) *Cache {
	c := &Cache{
		dataMap: make(map[string]cacheEntry),
	}
	go c.reapLoop(secInterval)
	return c
}
