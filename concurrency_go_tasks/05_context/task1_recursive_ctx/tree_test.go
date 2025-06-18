package tree

import (
	"context"
	"testing"
	"time"
)

func TestWalkCancel(t *testing.T) {
	root := &Node{Value: 1, Children: []*Node{{Value: 2}, {Value: 3}}}
	ctx, cancel := context.WithCancel(context.Background())
	visited := make([]int, 0)
	go func() {
		time.Sleep(20 * time.Millisecond)
		cancel()
	}()
	err := Walk(ctx, root, func(v int) {
		visited = append(visited, v)
		time.Sleep(15 * time.Millisecond)
	})
	if err == nil {
		t.Fatal("expected context cancellation")
	}
	if len(visited) >= 3 {
		t.Fatalf("walk should have been cancelled early, visited %v", visited)
	}
}

func TestWalkFull(t *testing.T) {
	root := &Node{Value: 1, Children: []*Node{{Value: 2}, {Value: 3}}}
	visited := make([]int, 0)
	err := Walk(context.Background(), root, func(v int) { visited = append(visited, v) })
	if err != nil {
		t.Fatalf("не ожидали ошибку: %v", err)
	}
	if len(visited) != 3 {
		t.Fatalf("ожидали пройти 3 узла, прошли %d", len(visited))
	}
}
