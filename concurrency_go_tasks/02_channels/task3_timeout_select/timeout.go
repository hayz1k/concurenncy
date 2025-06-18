package timeout

import (
	"context"
	"errors"
	"time"
)

func Work(ctx context.Context) error {
	done := make(chan struct{})
	go func() {
		time.Sleep(150 * time.Millisecond)
		close(done)
	}()

	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(100 * time.Millisecond):
		return errors.New("timeout")
	}
}
