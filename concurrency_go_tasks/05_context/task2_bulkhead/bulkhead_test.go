package bulkhead

import (
	"context"
	"sync/atomic"
	"testing"
	"time"
)

func TestBulkheadLimit(t *testing.T) {
	b := New(2)
	var running int32
	jobs := []func() error{
		func() error {
			atomic.AddInt32(&running, 1)
			time.Sleep(30 * time.Millisecond)
			atomic.AddInt32(&running, -1)
			return nil
		},
		func() error {
			atomic.AddInt32(&running, 1)
			time.Sleep(30 * time.Millisecond)
			atomic.AddInt32(&running, -1)
			return nil
		},
		func() error {
			atomic.AddInt32(&running, 1)
			time.Sleep(30 * time.Millisecond)
			atomic.AddInt32(&running, -1)
			return nil
		},
	}
	errCh := make(chan error, len(jobs))
	for _, job := range jobs {
		go func(j func() error) {
			errCh <- b.Do(context.Background(), j)
		}(job)
	}
	for range jobs {
		<-errCh
	}
	if running > 2 {
		t.Fatalf("running too many: %d", running)
	}
}

func TestBulkheadCancel(t *testing.T) {
	b := New(1)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err := b.Do(ctx, func() error { time.Sleep(10 * time.Millisecond); return nil })
	if err == nil {
		t.Fatal("expected cancellation error")
	}
}

func TestBulkheadSuccess(t *testing.T) {
	b := New(1)
	err := b.Do(context.Background(), func() error { return nil })
	if err != nil {
		t.Fatalf("не ожидали ошибку: %v", err)
	}
}
