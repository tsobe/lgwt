package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("Incrementing 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCount(t, counter, 3)
	})

	t.Run("Increment runs safely concurrently", func(t *testing.T) {
		wantCount := 1000
		counter := NewCounter()
		var wg sync.WaitGroup
		wg.Add(wantCount)

		for i := 0; i < wantCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCount(t, counter, wantCount)
	})
}

func assertCount(t *testing.T, counter *Counter, want int) {
	if counter.Value() != want {
		t.Errorf("Got %d, expected %d", counter.Value(), want)
	}
}
