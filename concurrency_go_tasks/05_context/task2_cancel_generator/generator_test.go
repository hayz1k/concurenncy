package generator

import (
	"context"
	"testing"
)

func TestGenerateNumbers(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := Generate(ctx)
	got := make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		n, ok := <-ch
		if !ok {
			t.Fatalf("ожидали %d элементов, получили %d", 5, len(got))
		}
		got = append(got, n)
	}
	for i, v := range got {
		if v != i {
			t.Fatalf("ожидали %d, получили %d", i, v)
		}
	}
}

func TestGenerateCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	ch := Generate(ctx)
	if _, ok := <-ch; ok {
		t.Fatal("канал должен закрыться после отмены")
	}
}
