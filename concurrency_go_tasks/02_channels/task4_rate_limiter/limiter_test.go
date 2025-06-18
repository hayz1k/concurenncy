package limiter

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestLimiter(t *testing.T) {
	l := NewLimiter()
	var allowed int32
	for i := 0; i < 10; i++ {
		if l.Allow() {
			atomic.AddInt32(&allowed, 1)
		}
	}
	if allowed > 5 {
		t.Fatalf("allowed more than 5 events: %d", allowed)
	}
	time.Sleep(time.Second)
	for i := 0; i < 5; i++ {
		if !l.Allow() {
			t.Fatalf("expected allowed after refill")
		}
	}
}
