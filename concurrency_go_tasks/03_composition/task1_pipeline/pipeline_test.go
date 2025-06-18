package pipeline

import "testing"

func TestRun(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	result := Run(nums)
	if result != 110 {
		t.Fatalf("expected 110, got %d", result)
	}
}
