package counter

import "sync"

type Counter struct {
	mu sync.Mutex
	v  int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	c.v++
	c.mu.Unlock()
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v
}
