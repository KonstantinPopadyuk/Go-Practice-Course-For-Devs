# 19. Concurrent JSONL Log Analysis

**Time:** 40–45 minutes.

Implement `Analyze(ctx, paths, workers)`: process multiple JSONL files with a bounded number of workers and aggregate their statistics.

Each non-empty line has the form `{"level":"INFO","latency_ms":12}`.

## Result and Contract

- `Report.Lines` is the number of valid, non-empty records.
- `Report.ByLevel` is the number of records at each level.
- `Report.MaxLatencyMS` is the maximum latency.
- If `workers <= 0`, return `ErrInvalidWorkers` before starting any goroutines or opening any files.
- No more than `workers` files may be open or processed concurrently.
- Read each file as a stream; support lines up to 1 MiB by configuring `Scanner.Buffer`.
- Ignore empty lines.
- Invalid JSON, an empty `level`, or a negative latency stops processing; the error must include the path and line number.
- On any error, return a zero-value `Report`, not a partial result that depends on completion order.
- Context cancellation stops dispatching new work and interrupts an active scan between lines.
- The final result must not depend on worker completion order and must pass with `-race`.

Run: `go test ./tasks/19-log-analyzer -race -v`.

A useful approach is to build a local `Report` for each file and merge the reports in one goroutine instead of sharing a map protected by a mutex.

## Going Further

Compute p50 and p95 latency with bounded memory, or explain which streaming algorithm you would choose.
