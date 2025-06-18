package debounce

import (
	"testing"
	"time"
)

func TestDebounce(t *testing.T) {
	in := make(chan int)
	out := Debounce(50*time.Millisecond, in)

	go func() {
		in <- 1
		time.Sleep(10 * time.Millisecond)
		in <- 2
		time.Sleep(10 * time.Millisecond)
		in <- 3
		close(in)
	}()

	v, ok := <-out
	if !ok || v != 3 {
		t.Fatalf("expected 3, got %d", v)
	}
}

func TestDebounceMultiple(t *testing.T) {
	in := make(chan int)
	out := Debounce(20*time.Millisecond, in)
	go func() {
		in <- 1
		time.Sleep(25 * time.Millisecond)
		in <- 2
		close(in)
	}()
	first, ok := <-out
	if !ok || first != 1 {
		t.Fatalf("ожидалось первое значение 1, получено %d", first)
	}
	second, ok := <-out
	if !ok || second != 2 {
		t.Fatalf("ожидалось второе значение 2, получено %d", second)
	}
}
