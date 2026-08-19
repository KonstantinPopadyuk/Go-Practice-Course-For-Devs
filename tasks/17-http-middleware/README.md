# 17. HTTP Middleware Without a Framework

**Time:** 25–30 minutes.

Implement `Middleware(logger, next)`, combining request IDs, access logging, and panic recovery for `net/http`.

## Contract

- If the request contains `X-Request-ID`, preserve it; otherwise, generate a non-empty unique ID safely under concurrent requests.
- Add the ID to the response header and the context; `RequestID(ctx)` makes it available to the handler.
- Recover from panics and log them without crashing the process. Return 500 if the response has not started; after `Write` or `WriteHeader`, preserving the status already sent is the only possible behavior.
- The access log must include the method, path, status, request ID, and duration.
- Preserve the status of a normal handler; if it calls `Write` without `WriteHeader`, the status is 200.
- Do not log the query string or sensitive headers.

Run: `go test ./tasks/17-http-middleware -race -v`.

You will need a wrapper around `http.ResponseWriter`. Do not add methods such as `http.Flusher` blindly: production middleware should delegate only the optional interfaces supported by the underlying writer.

## Going Further

Correctly support `Flusher`, `Hijacker`, and `Pusher`, and produce structured logs with `log/slog`.
