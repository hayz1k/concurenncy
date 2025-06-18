package cache

import (
	"sync"
	"testing"
)

func TestCache(t *testing.T) {
	c := New()
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		i := i
		wg.Add(1)
		go func() {
			c.Set("key", i)
			wg.Done()
		}()
	}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			c.Get("key")
			wg.Done()
		}()
	}
	wg.Wait()
	if _, ok := c.Get("key"); !ok {
		t.Fatal("expected key to exist")
	}
}
