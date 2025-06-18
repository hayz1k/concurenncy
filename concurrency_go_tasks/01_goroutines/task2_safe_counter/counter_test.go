package counter

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	var c Counter
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 10; j++ {
				c.Inc()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	if v := c.Value(); v != 1000 {
		t.Fatalf("expected 1000, got %d", v)
	}
}

func TestCounterInitialValue(t *testing.T) {
	var c Counter
	if v := c.Value(); v != 0 {
		t.Fatalf("новый счётчик должен быть 0, получено %d", v)
	}
}
