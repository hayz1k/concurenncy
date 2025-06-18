package cache

import "sync"

type Cache struct {
	mu   sync.RWMutex
	data map[string]interface{}
}

func New() *Cache {
	return &Cache{data: make(map[string]interface{})}
}

func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	c.data[key] = value
	c.mu.Unlock()
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	v, ok := c.data[key]
	c.mu.RUnlock()
	return v, ok
}
