# 15. Deduplicating Concurrent Requests

**Time:** 40–45 minutes.

Implement `Group[K, V].Do`: concurrent requests for the same key must share a single call to an expensive function. This is a simplified version of `singleflight`.

## Contract

- The first caller for a key starts `fn`; all other callers wait for the same result.
- Functions for different keys may run concurrently.
- The entry is removed after completion, so a later call starts `fn` again.
- A waiting caller may return `ctx.Err()` without canceling the shared work or the other waiting callers.
- An error from `fn` is shared in the same way as a successful result.
- `fn` must not be called while holding the mutex.
- The type must pass `go test -race`.

Run: `go test ./tasks/15-singleflight -race -v`.

A useful design is a map from each key to a call structure containing a `done` channel, a result, and an error. Closing the channel publishes the previously written fields to all waiting callers.

## Going Further

Return a third `shared` flag, add `Forget(key)`, and define the behavior when `fn` panics.
