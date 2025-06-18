package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestPingPong(t *testing.T) {
	var buf bytes.Buffer
	PingPong(&buf)
	output := strings.Split(strings.TrimSpace(buf.String()), "\n")
	if len(output) != 10 {
		t.Fatalf("expected 10 lines, got %d", len(output))
	}
	pingCount, pongCount := 0, 0
	for i, line := range output {
		if i%2 == 0 && line != "ping" {
			t.Errorf("expected ping at line %d, got %s", i, line)
		}
		if i%2 == 1 && line != "pong" {
			t.Errorf("expected pong at line %d, got %s", i, line)
		}
		if line == "ping" {
			pingCount++
		} else if line == "pong" {
			pongCount++
		}
	}
	if pingCount != 5 || pongCount != 5 {
		t.Fatalf("expected 5 ping and 5 pong, got %d ping and %d pong", pingCount, pongCount)
	}
}

func TestPingPongStartsWithPing(t *testing.T) {
	var buf bytes.Buffer
	PingPong(&buf)
	output := strings.Split(strings.TrimSpace(buf.String()), "\n")
	if len(output) == 0 || output[0] != "ping" {
		t.Fatal("первой строкой должен быть ping")
	}
}
