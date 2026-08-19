# 13. The First Successful Result

**Time:** 30–35 minutes.

Implement `FirstSuccess`: run several equivalent operations concurrently (for example, requests to mirrors) and return the first successful result.

## Contract

- Start all functions concurrently with a child context.
- Return the first result with a `nil` error immediately and cancel the remaining operations.
- If every operation fails, return `errors.Join` of those errors so that `errors.Is` can find each cause.
- An empty list returns `ErrNoOperations`.
- If the parent context is canceled, return its error and stop the child operations.
- No sender may remain blocked forever after an early return. Consider the channel buffer size.

Run: `go test ./tasks/13-first-success -race -v`.

This is a simplified model of hedged requests: duplicating a request can sometimes reduce tail latency, but it should be used sparingly in a production service.

## Going Further

Add a delay between operation starts and a metric that counts canceled attempts.
