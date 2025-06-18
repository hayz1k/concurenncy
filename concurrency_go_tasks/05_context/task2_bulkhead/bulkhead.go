package bulkhead

import (
	"context"
)

type Bulkhead struct {
	sem chan struct{}
}

func New(limit int) *Bulkhead {
	return &Bulkhead{sem: make(chan struct{}, limit)}
}

func (b *Bulkhead) Do(ctx context.Context, fn func() error) error {
	select {
	case b.sem <- struct{}{}:
		defer func() { <-b.sem }()
	case <-ctx.Done():
		return ctx.Err()
	}

	done := make(chan error, 1)
	go func() {
		done <- fn()
	}()

	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
