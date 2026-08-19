# 10. A Robust HTTP Client

**Time:** 25–30 minutes.

Implement `FetchJSON(ctx, client, url, maxBytes, dst)`: perform a GET request and decode the JSON safely for the calling service.

## Contract

- Create the request with `http.NewRequestWithContext`.
- For a 2xx response, decode exactly one JSON value into `dst`; trailing JSON or other data is an error.
- A response larger than `maxBytes` must return `ErrResponseTooLarge` (checked with `errors.Is`).
- For a non-2xx response, return an error containing the status code and a small excerpt from the body (no more than 4 KiB).
- Always close the response body; wrap transport errors.
- Do not create an `http.Client` inside the function: accepting a client makes the code testable and preserves its connection pool.
- A canceled context must interrupt the request promptly.

Run: `go test ./tasks/10-http-client -v`.

Limit hint: read at most `maxBytes+1` bytes. Using only `io.LimitReader(..., maxBytes)` is not enough to distinguish a response of exactly the permitted size from a truncated response.

## Going Further

Add retries only for idempotent requests and 429/5xx responses, using exponential backoff that respects the context.
