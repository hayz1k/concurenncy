package pool

import (
	"testing"
	"time"
)

func TestRunPool(t *testing.T) {
	jobs := []int{1, 2, 3, 4, 5, 6}
	start := time.Now()
	sum := RunPool(jobs, 3)
	duration := time.Since(start)
	if sum != 21 {
		t.Fatalf("expected 21, got %d", sum)
	}
	if duration >= time.Duration(len(jobs))*10*time.Millisecond {
		t.Fatal("expected parallel execution")
	}
}
