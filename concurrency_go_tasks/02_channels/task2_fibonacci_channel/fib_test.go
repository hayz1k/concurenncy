package fibonacci

import "testing"

func TestFib(t *testing.T) {
	ch := Fib(6)
	expected := []int{0, 1, 1, 2, 3, 5}
	i := 0
	for v := range ch {
		if v != expected[i] {
			t.Fatalf("expected %d, got %d", expected[i], v)
		}
		i++
	}
	if i != len(expected) {
		t.Fatalf("expected %d numbers, got %d", len(expected), i)
	}
}
