package ch02

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	var counter int64
	var wg sync.WaitGroup
	var mu sync.Mutex
	for range 64 {
		wg.Add(1)
		go func() {
			for range 1000000 {
				mu.Lock()
				counter++
				mu.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	if counter != 64000000 {
		t.Errorf("Expected 64000000, got %d", counter)
	}
}