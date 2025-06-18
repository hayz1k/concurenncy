package scheduler

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestEvery(t *testing.T) {
	var count int32
	stop := Every(50*time.Millisecond, func() { atomic.AddInt32(&count, 1) })
	time.Sleep(170 * time.Millisecond)
	stop()
	if c := atomic.LoadInt32(&count); c < 3 || c > 4 {
		t.Fatalf("expected 3 or 4 executions, got %d", c)
	}
}

func TestEveryStop(t *testing.T) {
	var count int32
	stop := Every(10*time.Millisecond, func() { atomic.AddInt32(&count, 1) })
	time.Sleep(25 * time.Millisecond)
	stop()
	c := atomic.LoadInt32(&count)
	time.Sleep(30 * time.Millisecond)
	if atomic.LoadInt32(&count) != c {
		t.Fatal("функция должна быть остановлена")
	}
}
