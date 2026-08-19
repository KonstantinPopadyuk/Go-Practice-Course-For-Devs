package task16

import (
	"context"
	"errors"
	"sort"
	"sync"
	"testing"
	"time"
)

func TestLimiterSpacesConcurrentCalls(t *testing.T) {
	const interval = 15 * time.Millisecond
	l, err := NewLimiter(interval)
	if err != nil {
		t.Fatal(err)
	}
	times := make([]time.Time, 3)
	var wg sync.WaitGroup
	for i := range times {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if err := l.Wait(context.Background()); err != nil {
				t.Errorf("Wait: %v", err)
			}
			times[i] = time.Now()
		}(i)
	}
	wg.Wait()
	sort.Slice(times, func(i, j int) bool { return times[i].Before(times[j]) })
	const schedulingTolerance = 3 * time.Millisecond
	for i := 1; i < len(times); i++ {
		if gap := times[i].Sub(times[i-1]); gap < interval-schedulingTolerance {
			t.Fatalf("slots %d and %d are only %v apart", i-1, i, gap)
		}
	}
}

func TestLimiterValidationAndCancellation(t *testing.T) {
	if _, err := NewLimiter(0); !errors.Is(err, ErrInvalidInterval) {
		t.Fatalf("NewLimiter error = %v", err)
	}
	if _, err := NewLimiter(-time.Second); !errors.Is(err, ErrInvalidInterval) {
		t.Fatalf("NewLimiter negative interval error = %v", err)
	}
	l, _ := NewLimiter(time.Hour)
	if err := l.Wait(context.Background()); err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := l.Wait(ctx); !errors.Is(err, context.Canceled) {
		t.Fatalf("Wait error = %v", err)
	}
}
