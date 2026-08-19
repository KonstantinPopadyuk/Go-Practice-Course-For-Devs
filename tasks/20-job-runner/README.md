# 20. Mini Job Scheduler

**Time:** 45–60 minutes.

This final task combines a worker pool, context cancellation, timeouts, retries, panic recovery, and deterministic results. Implement `RunJobs`.

```go
type Job struct {
    ID         string
    MaxRetries int
    Run        func(context.Context) error
}
```

## Contract

- When work is available, use up to `workers` worker goroutines, but never run more than `workers` jobs concurrently; results must preserve input order.
- Each attempt receives a child context with `attemptTimeout`.
- `MaxRetries` is the number of retries **after** the first attempt. Treat a negative value as 0.
- Retry any error until all attempts are exhausted; check the parent context between attempts.
- Success stops further retries. `Result.Attempts` contains the actual number of times the job was started.
- A panic inside `Job.Run` becomes an error containing the job ID, and the worker continues running.
- When the parent context is canceled, do not dispatch new jobs. For jobs that were not started, set `Attempts == 0` and `ErrNotStarted`; for jobs already started, `Result.Err` must be compatible with `ctx.Err()`.
- An empty ID, a nil function, `workers <= 0`, or a non-positive timeout is a configuration error that must be detected before any job starts.
- The second return value, `error`, is reserved for configuration errors. After successful validation, execution errors, timeouts, panics, and parent-context cancellation are recorded in the corresponding `Result` values, which `RunJobs` returns together with `nil`.
- All goroutines must terminate; assume that `Job.Run` respects the context passed to it.

Run: `go test ./tasks/20-job-runner -race -v`.

Split the solution into small pieces: validation, one safe attempt with recovery, a retry loop, and a worker-pool dispatcher. Do not try to implement everything in one function.

## Going Further

Add exponential backoff with jitter, retryable-error classification, and a progress callback. Then expose the runner through an HTTP API and apply the middleware from task 17.
