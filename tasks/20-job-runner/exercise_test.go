package task20

import (
	"context"
	"errors"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

func TestRunJobsRetriesAndPreservesOrder(t *testing.T) {
	var flakyCalls int32
	jobs := []Job{
		{ID: "slow", Run: func(ctx context.Context) error { time.Sleep(15 * time.Millisecond); return nil }},
		{ID: "flaky", MaxRetries: 2, Run: func(ctx context.Context) error {
			if atomic.AddInt32(&flakyCalls, 1) < 3 {
				return errors.New("temporary")
			}
			return nil
		}},
		{ID: "fast", Run: func(ctx context.Context) error { return nil }},
	}
	got, err := RunJobs(context.Background(), jobs, 2, time.Second)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 3 || got[0].ID != "slow" || got[1].ID != "flaky" || got[2].ID != "fast" {
		t.Fatalf("result order = %+v", got)
	}
	if got[1].Attempts != 3 || got[1].Err != nil {
		t.Fatalf("flaky result = %+v", got[1])
	}
}

func TestRunJobsTimeoutAndPanic(t *testing.T) {
	var afterPanicCalls int32
	jobs := []Job{
		{ID: "timeout", Run: func(ctx context.Context) error { <-ctx.Done(); return ctx.Err() }},
		{ID: "crasher", Run: func(context.Context) error { panic("boom") }},
		{ID: "after", Run: func(context.Context) error { atomic.AddInt32(&afterPanicCalls, 1); return nil }},
	}
	got, err := RunJobs(context.Background(), jobs, 1, 10*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != len(jobs) {
		t.Fatalf("got %d results, want %d", len(got), len(jobs))
	}
	if !errors.Is(got[0].Err, context.DeadlineExceeded) {
		t.Fatalf("timeout result = %+v", got[0])
	}
	if got[1].Err == nil || !strings.Contains(got[1].Err.Error(), "panic") || !strings.Contains(got[1].Err.Error(), "crasher") || !strings.Contains(got[1].Err.Error(), "boom") {
		t.Fatalf("panic result = %+v", got[1])
	}
	if got[2].Err != nil || got[2].Attempts != 1 || atomic.LoadInt32(&afterPanicCalls) != 1 {
		t.Fatalf("worker did not continue after panic: result = %+v, calls = %d", got[2], afterPanicCalls)
	}
}

func TestRunJobsValidatesBeforeStarting(t *testing.T) {
	var calls int32
	jobs := []Job{
		{ID: "valid", Run: func(context.Context) error { atomic.AddInt32(&calls, 1); return nil }},
		{ID: "", Run: func(context.Context) error { return nil }},
	}
	if _, err := RunJobs(context.Background(), jobs, 1, time.Second); err == nil {
		t.Fatal("expected validation error")
	}
	if calls != 0 {
		t.Fatal("job ran before validation completed")
	}
}

func TestRunJobsRejectsInvalidConfiguration(t *testing.T) {
	job := Job{ID: "valid", Run: func(context.Context) error { return nil }}
	if _, err := RunJobs(context.Background(), []Job{job}, 0, time.Second); err == nil {
		t.Fatal("expected workers validation error")
	}
	if _, err := RunJobs(context.Background(), []Job{job}, 1, 0); err == nil {
		t.Fatal("expected attempt timeout validation error")
	}
	if _, err := RunJobs(context.Background(), []Job{{ID: "nil-run"}}, 1, time.Second); err == nil {
		t.Fatal("expected nil Run validation error")
	}
}

func TestRunJobsCancellationMarksNotStarted(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	started := make(chan struct{})
	var waitingCalls int32
	jobs := []Job{
		{ID: "running", Run: func(ctx context.Context) error {
			close(started)
			<-ctx.Done()
			return ctx.Err()
		}},
		{ID: "waiting", Run: func(context.Context) error {
			atomic.AddInt32(&waitingCalls, 1)
			return nil
		}},
	}
	type outcome struct {
		results []Result
		err     error
	}
	done := make(chan outcome, 1)
	go func() {
		results, err := RunJobs(ctx, jobs, 1, time.Second)
		done <- outcome{results: results, err: err}
	}()
	select {
	case <-started:
	case <-time.After(time.Second):
		t.Fatal("first job did not start")
	}
	cancel()
	select {
	case got := <-done:
		if got.err != nil {
			t.Fatalf("RunJobs returned execution error: %v", got.err)
		}
		if len(got.results) != 2 {
			t.Fatalf("got %d results, want 2", len(got.results))
		}
		if got.results[0].Attempts != 1 || !errors.Is(got.results[0].Err, context.Canceled) {
			t.Fatalf("running result = %+v", got.results[0])
		}
		if got.results[1].Attempts != 0 || !errors.Is(got.results[1].Err, ErrNotStarted) {
			t.Fatalf("not-started result = %+v", got.results[1])
		}
		if atomic.LoadInt32(&waitingCalls) != 0 {
			t.Fatal("job started after parent cancellation")
		}
	case <-time.After(time.Second):
		t.Fatal("RunJobs did not return after parent cancellation")
	}
}

func TestRunJobsUsesBoundedConcurrency(t *testing.T) {
	var active, peak int32
	started := make(chan struct{}, 3)
	release := make(chan struct{})
	jobs := make([]Job, 3)
	for i := range jobs {
		jobs[i] = Job{ID: string(rune('a' + i)), Run: func(context.Context) error {
			current := atomic.AddInt32(&active, 1)
			defer atomic.AddInt32(&active, -1)
			for {
				old := atomic.LoadInt32(&peak)
				if current <= old || atomic.CompareAndSwapInt32(&peak, old, current) {
					break
				}
			}
			started <- struct{}{}
			<-release
			return nil
		}}
	}
	type outcome struct {
		results []Result
		err     error
	}
	done := make(chan outcome, 1)
	go func() {
		results, err := RunJobs(context.Background(), jobs, 2, time.Second)
		done <- outcome{results: results, err: err}
	}()
	for range 2 {
		select {
		case <-started:
		case <-time.After(time.Second):
			close(release)
			t.Fatal("RunJobs did not use the available workers")
		}
	}
	select {
	case <-started:
		close(release)
		t.Fatal("RunJobs exceeded the worker limit")
	case <-time.After(25 * time.Millisecond):
	}
	close(release)
	select {
	case got := <-done:
		if got.err != nil || len(got.results) != len(jobs) {
			t.Fatalf("RunJobs = (%+v, %v)", got.results, got.err)
		}
		for _, result := range got.results {
			if result.Err != nil {
				t.Fatalf("job result = %+v", result)
			}
		}
		if peak != 2 {
			t.Fatalf("peak concurrency = %d, want 2", peak)
		}
	case <-time.After(time.Second):
		t.Fatal("RunJobs did not finish")
	}
}
