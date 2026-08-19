package task11

import (
	"context"
	"errors"
	"reflect"
	"regexp"
	"sync/atomic"
	"testing"
	"time"
)

func TestParallelMapOrderAndLimit(t *testing.T) {
	var active, peak int32
	fn := func(ctx context.Context, n int) (int, error) {
		current := atomic.AddInt32(&active, 1)
		defer atomic.AddInt32(&active, -1)
		for {
			old := atomic.LoadInt32(&peak)
			if current <= old || atomic.CompareAndSwapInt32(&peak, old, current) {
				break
			}
		}
		time.Sleep(time.Duration(6-n) * time.Millisecond)
		return n * n, nil
	}
	got, err := ParallelMap(context.Background(), 2, []int{1, 2, 3, 4, 5}, fn)
	if err != nil {
		t.Fatal(err)
	}
	if want := []int{1, 4, 9, 16, 25}; !reflect.DeepEqual(got, want) {
		t.Fatalf("results = %v, want %v", got, want)
	}
	if peak > 2 {
		t.Fatalf("peak concurrency = %d, want <= 2", peak)
	}
}

func TestParallelMapErrorAndCancellation(t *testing.T) {
	boom := errors.New("boom")
	firstStarted := make(chan struct{})
	firstStopped := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	done := make(chan error, 1)
	go func() {
		_, err := ParallelMap(ctx, 2, []int{10, 20}, func(ctx context.Context, n int) (int, error) {
			if n == 20 {
				<-firstStarted
				return 0, boom
			}
			close(firstStarted)
			<-ctx.Done()
			close(firstStopped)
			return 0, ctx.Err()
		})
		done <- err
	}()

	var err error
	select {
	case err = <-done:
	case <-time.After(time.Second):
		cancel()
		t.Fatal("ParallelMap did not cancel in-flight work after an error")
	}
	if !errors.Is(err, boom) {
		t.Fatalf("error = %v, want wrapped boom", err)
	}
	if !regexp.MustCompile(`\bindex 1\b`).MatchString(err.Error()) {
		t.Fatalf("error = %q, want zero-based index 1", err)
	}
	select {
	case <-firstStopped:
	case <-time.After(time.Second):
		t.Fatal("in-flight callback did not observe cancellation")
	}
	if _, err := ParallelMap(context.Background(), 0, nil, nil); !errors.Is(err, ErrInvalidWorkers) {
		t.Fatalf("workers error = %v", err)
	}
}

func TestParallelMapParentCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	started := make(chan struct{}, 2)
	stopped := make(chan struct{}, 2)
	done := make(chan error, 1)
	go func() {
		_, err := ParallelMap(ctx, 2, []int{1, 2}, func(ctx context.Context, n int) (int, error) {
			started <- struct{}{}
			<-ctx.Done()
			stopped <- struct{}{}
			return 0, ctx.Err()
		})
		done <- err
	}()
	for range 2 {
		select {
		case <-started:
		case <-time.After(time.Second):
			t.Fatal("callbacks did not start")
		}
	}
	cancel()
	select {
	case err := <-done:
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("error = %v, want context.Canceled", err)
		}
	case <-time.After(time.Second):
		t.Fatal("ParallelMap did not return after parent cancellation")
	}
	for range 2 {
		select {
		case <-stopped:
		case <-time.After(time.Second):
			t.Fatal("callback did not stop")
		}
	}
}
