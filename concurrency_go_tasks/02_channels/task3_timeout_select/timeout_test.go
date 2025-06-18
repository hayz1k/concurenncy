package timeout

import (
	"context"
	"testing"
)

func TestWorkTimeout(t *testing.T) {
	err := Work(context.Background())
	if err == nil {
		t.Fatal("expected timeout error")
	}
}

func TestWorkCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := Work(ctx); err == nil {
		t.Fatal("ожидалась ошибка отмены контекста")
	}
}
