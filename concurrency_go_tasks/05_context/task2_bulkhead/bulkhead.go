package bulkhead

import (
	"context"
)

// Bulkhead ограничивает количество одновременных вызовов.
type Bulkhead struct {
	sem chan struct{}
}

// New создаёт Bulkhead с ограничением limit.
func New(limit int) *Bulkhead {
	// TODO: инициализировать семафор с указанным лимитом
	return &Bulkhead{}
}

// Do выполняет fn с учётом ограничения и поддержкой отмены через ctx.
func (b *Bulkhead) Do(ctx context.Context, fn func() error) error {
	// TODO: реализовать выполнение функции с использованием семафора
	return nil
}
