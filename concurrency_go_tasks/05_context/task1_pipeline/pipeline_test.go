package pipelinectx

import (
	"context"
	"testing"
)

func TestRun(t *testing.T) {
	result, err := Run(context.Background(), []int{1, 2, 3, 4, 5})
	if err != nil {
		t.Fatalf("не ожидали ошибку: %v", err)
	}
	if result != 30 {
		t.Fatalf("ожидали 30, получили %d", result)
	}
}

func TestRunCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := Run(ctx, []int{1, 2}); err == nil {
		t.Fatal("ожидалась ошибка отмены")
	}
}

func TestRunEmpty(t *testing.T) {
	res, err := Run(context.Background(), nil)
	if err != nil {
		t.Fatalf("не ожидали ошибку: %v", err)
	}
	if res != 0 {
		t.Fatalf("при пустом списке результат должен быть 0")
	}
}
