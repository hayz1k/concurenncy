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
