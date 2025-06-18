package producerconsumer

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestProducerConsumer(t *testing.T) {
	var buf bytes.Buffer
	Run(&buf)
	output := strings.Split(strings.TrimSpace(buf.String()), "\n")
	if len(output) != 10 {
		t.Fatalf("expected 10 numbers, got %d", len(output))
	}
	for i, line := range output {
		expected := i + 1
		if line != fmt.Sprint(expected) {
			t.Errorf("expected %d at line %d, got %s", expected, i, line)
		}
	}
}

func TestProducerConsumerFirstLast(t *testing.T) {
	var buf bytes.Buffer
	Run(&buf)
	output := strings.Split(strings.TrimSpace(buf.String()), "\n")
	if len(output) > 0 {
		if output[0] != "1" || output[len(output)-1] != "10" {
			t.Fatal("последовательность должна начинаться с 1 и заканчиваться 10")
		}
	}
}
