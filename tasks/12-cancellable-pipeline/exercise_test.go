package task12

import (
	"context"
	"reflect"
	"testing"
	"time"
)

func TestPipeline(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	out := Square(ctx, Generate(ctx, 1, 2, 3, 4, 5))
	got := Take(ctx, out, 3)
	cancel()
	if want := []int{1, 4, 9}; !reflect.DeepEqual(got, want) {
		t.Fatalf("Take = %v, want %v", got, want)
	}
}

func TestSquareStopsOnCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	in := make(chan int)
	out := Square(ctx, in)
	cancel()
	select {
	case _, ok := <-out:
		if ok {
			t.Fatal("Square produced a value without input")
		}
	case <-time.After(time.Second):
		t.Fatal("Square stayed blocked receiving input after cancellation")
	}
}

func TestTakeStopsOnCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	done := make(chan []int, 1)
	go func() { done <- Take(ctx, make(chan int), 1) }()
	select {
	case got := <-done:
		if len(got) != 0 {
			t.Fatalf("Take = %v, want empty result", got)
		}
	case <-time.After(time.Second):
		t.Fatal("Take did not return after cancellation")
	}
}

func TestTakeEdgeCases(t *testing.T) {
	closed := make(chan int)
	close(closed)
	if got := Take(context.Background(), closed, 5); len(got) != 0 {
		t.Fatalf("Take(closed) = %v", got)
	}
	if got := Take(context.Background(), closed, 0); len(got) != 0 {
		t.Fatalf("Take(n=0) = %v", got)
	}
}
