package task13

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestFirstSuccess(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	slowCanceled := make(chan struct{})
	type result struct {
		value string
		err   error
	}
	done := make(chan result, 1)
	go func() {
		got, err := FirstSuccess(ctx,
			func(ctx context.Context) (string, error) {
				<-ctx.Done()
				close(slowCanceled)
				return "", ctx.Err()
			},
			func(ctx context.Context) (string, error) { return "mirror-b", nil },
		)
		done <- result{value: got, err: err}
	}()
	var got string
	var err error
	select {
	case result := <-done:
		got, err = result.value, result.err
	case <-time.After(time.Second):
		cancel()
		t.Fatal("operations did not run concurrently")
	}
	if err != nil || got != "mirror-b" {
		t.Fatalf("FirstSuccess = (%q, %v)", got, err)
	}
	select {
	case <-slowCanceled:
	case <-time.After(100 * time.Millisecond):
		t.Fatal("losing operation was not canceled")
	}
}

func TestFirstSuccessParentCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	started := make(chan struct{}, 2)
	stopped := make(chan struct{}, 2)
	op := func(ctx context.Context) (string, error) {
		started <- struct{}{}
		<-ctx.Done()
		stopped <- struct{}{}
		return "", ctx.Err()
	}
	done := make(chan error, 1)
	go func() {
		_, err := FirstSuccess(ctx, op, op)
		done <- err
	}()
	for range 2 {
		select {
		case <-started:
		case <-time.After(time.Second):
			t.Fatal("operations did not start")
		}
	}
	cancel()
	select {
	case err := <-done:
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("error = %v, want context.Canceled", err)
		}
	case <-time.After(time.Second):
		t.Fatal("FirstSuccess did not return after parent cancellation")
	}
	for range 2 {
		select {
		case <-stopped:
		case <-time.After(time.Second):
			t.Fatal("operation did not stop")
		}
	}
}

func TestFirstSuccessJoinsErrors(t *testing.T) {
	errA, errB := errors.New("a"), errors.New("b")
	_, err := FirstSuccess(context.Background(),
		func(context.Context) (string, error) { return "", errA },
		func(context.Context) (string, error) { return "", errB },
	)
	if !errors.Is(err, errA) || !errors.Is(err, errB) {
		t.Fatalf("error = %v, want both causes", err)
	}
	if _, err := FirstSuccess(context.Background()); !errors.Is(err, ErrNoOperations) {
		t.Fatalf("empty error = %v", err)
	}
}
