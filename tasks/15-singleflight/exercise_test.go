package task15

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"
	"time"
)

func TestGroupSharesInFlightCall(t *testing.T) {
	var g Group[string, int]
	var calls int32
	started := make(chan struct{})
	release := make(chan struct{})
	leaderDone := make(chan int, 1)
	go func() {
		value, _ := g.Do(context.Background(), "key", func() (int, error) {
			atomic.AddInt32(&calls, 1)
			close(started)
			<-release
			return 42, nil
		})
		leaderDone <- value
	}()
	select {
	case <-started:
	case <-time.After(time.Second):
		t.Fatal("leader did not start")
	}
	time.AfterFunc(50*time.Millisecond, func() { close(release) })
	got, err := g.Do(context.Background(), "key", func() (int, error) {
		atomic.AddInt32(&calls, 1)
		return 99, errors.New("must not run")
	})
	if err != nil || got != 42 {
		t.Fatalf("waiter result = (%d, %v), want (42, nil)", got, err)
	}
	if leader := <-leaderDone; leader != 42 {
		t.Fatalf("leader result = %d, want 42", leader)
	}
	if got := atomic.LoadInt32(&calls); got != 1 {
		t.Fatalf("fn called %d times, want 1", got)
	}
}

func TestGroupWaiterCanCancel(t *testing.T) {
	var g Group[string, int]
	release := make(chan struct{})
	started := make(chan struct{})
	done := make(chan struct{})
	go func() {
		defer close(done)
		_, _ = g.Do(context.Background(), "key", func() (int, error) {
			close(started)
			<-release
			return 7, nil
		})
	}()
	select {
	case <-started:
	case <-time.After(time.Second):
		t.Fatal("leader did not start")
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := g.Do(ctx, "key", func() (int, error) { return 99, errors.New("must not run") }); !errors.Is(err, context.Canceled) {
		t.Fatalf("waiter error = %v", err)
	}
	close(release)
	<-done
}

func TestGroupDifferentKeysRunConcurrently(t *testing.T) {
	var g Group[string, int]
	started := make(chan string, 2)
	release := make(chan struct{})
	done := make(chan struct{}, 2)
	for _, key := range []string{"a", "b"} {
		key := key
		go func() {
			_, _ = g.Do(context.Background(), key, func() (int, error) {
				started <- key
				<-release
				return 1, nil
			})
			done <- struct{}{}
		}()
	}
	for range 2 {
		select {
		case <-started:
		case <-time.After(time.Second):
			close(release)
			t.Fatal("different keys were serialized")
		}
	}
	close(release)
	for range 2 {
		select {
		case <-done:
		case <-time.After(time.Second):
			t.Fatal("Do did not return")
		}
	}
}

func TestGroupRunsAgainAfterCompletion(t *testing.T) {
	var g Group[string, int]
	var calls int
	for want := 1; want <= 2; want++ {
		got, err := g.Do(context.Background(), "key", func() (int, error) {
			calls++
			return calls, nil
		})
		if err != nil || got != want {
			t.Fatalf("call %d = (%d, %v)", want, got, err)
		}
	}
}
