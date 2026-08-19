# 16. Context-Aware Rate Limiter

**Time:** 30–35 minutes.

Implement `Limiter`, which allows operations to start no more than once per configured `interval`. Multiple goroutines may call `Wait` concurrently.

## Contract

- `NewLimiter(interval)` rejects a non-positive interval with `ErrInvalidInterval`.
- The first `Wait` may return immediately.
- Subsequent grants must be separated by at least `interval`.
- Concurrent callers receive distinct time slots without data races.
- `Wait(ctx)` returns `ctx.Err()` if the context is canceled before the reserved slot.
- Do not use a busy loop or `time.Sleep`: wait on a timer and select between it and `ctx.Done()`.
- A canceled caller may leave its reserved slot unused.

Run: `go test ./tasks/16-rate-limiter -race -v`.

Store the time of the next available slot under a mutex. Reserve the slot quickly, release the lock, and only then wait; otherwise, one slow caller will block everyone else on the mutex.

## Going Further

Turn the limiter into a token bucket with burst capacity and an injectable clock for fully deterministic tests.
