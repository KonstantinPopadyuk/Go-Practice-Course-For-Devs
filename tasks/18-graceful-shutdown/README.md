# 18. HTTP Server Lifecycle

**Time:** 35–45 minutes.

Implement `Run(ctx, server, listener, shutdownTimeout)`. The function starts an HTTP server and shuts it down gracefully when the context is canceled.

## Contract

- `server.Serve(listener)` must run in a separate goroutine.
- If `Serve` returns a real error before cancellation, return it with contextual information. `http.ErrServerClosed` is not considered an error.
- After `ctx.Done()`, call `server.Shutdown` with a **new** context and a timeout: the original context has already been canceled.
- Wait for `Serve` to finish so that its goroutine does not leak.
- Do not discard an error from `Shutdown`.
- If `shutdownTimeout <= 0`, return `ErrInvalidShutdownTimeout` before starting the server.

Run: `go test ./tasks/18-graceful-shutdown -race -v`.

In a real `main`, the context is usually canceled on `SIGINT` or `SIGTERM` through `signal.NotifyContext`. Here, a regular context replaces the signal so that the lifecycle is easy to test.

## Going Further

Add a readiness flag that becomes false before shutdown begins, and force a `Close` if the graceful shutdown timeout expires.
