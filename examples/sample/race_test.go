package sample_test

import (
	"sync"
	"testing"
)

func TestConcurrentIncrement(t *testing.T) {
	var mu sync.Mutex
	count := 0

	var wg sync.WaitGroup
	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()
		}()
	}
	wg.Wait()

	if count != 10 {
		t.Fatalf("expected 10, got %d", count)
	}
}
