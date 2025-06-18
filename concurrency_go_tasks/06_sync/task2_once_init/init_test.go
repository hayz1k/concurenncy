package initonce

import (
	"sync"
	"testing"
)

func TestInitOnce(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			Init()
			wg.Done()
		}()
	}
	wg.Wait()
	if !Initialized() {
		t.Fatal("resource not initialized")
	}
}

func TestInitializedMultiple(t *testing.T) {
	Init()
	if !Initialized() || !Initialized() {
		t.Fatal("ожидалось истинное значение инициализации")
	}
}
