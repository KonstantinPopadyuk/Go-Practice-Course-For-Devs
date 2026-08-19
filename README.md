# Go Practice: 20 Hands-On Exercises

A compact course for developers with Python experience who already know Go syntax and want to build fluency with Go idioms and common real-world tasks.

## Repository Structure

- The repository contains 20 independent exercises in `tasks/01-...` through `tasks/20-...`.
- Each directory contains a detailed task description, starter code, and tests.
- The exercises progress from foundational idioms to concurrency and a small service.
- There are no external dependencies: only the standard library is used.
- The tests fail at first. This is expected: find the `TODO` in the relevant directory and implement the contract.
- The `main` branch always contains the starter code, while personal progress on the first exercises is preserved in the `my-solution` branch.

## Quick Start

Go 1.22 or later is required.

```bash
go test ./tasks/01-slices
go test ./tasks/01-slices -run TestUniqueStable -v
go test ./tasks/11-worker-pool -race
go test ./...
```

You can run the same checks through the `Makefile`:

```bash
make check
make task TASK=01-slices
make race TASK=11-worker-pool
```

`make check` runs `go vet` and verifies that every package and test compiles without completed solutions. `make test` runs all exercise checks, so it is expected to remain red on the starter branch.

Switch between the clean course and personal progress with:

```bash
git switch main
git switch my-solution
```

The completion checklist is in [`PROGRESS.md`](PROGRESS.md).

A useful workflow:

1. Read the `README.md` for the selected exercise.
2. Run its tests and inspect the first failure.
3. Implement the smallest solution that satisfies the contract.
4. Run `go test`, then `go test -race` for concurrency exercises.
5. Once the tests pass, try the extension in the "Going Further" section.

## Concurrency Interview Track

Exercises 11–20 form a separate track with steadily increasing difficulty. Nine of them directly cover goroutines, channels, synchronization, or cancellation; exercise 17 applies the same production-minded habits to HTTP middleware.

- **11–13 — fundamentals:** a bounded worker pool, a cancellable pipeline, and the first successful concurrent result.
- **14–16 — shared state:** a thread-safe TTL cache, single-flight request deduplication, and a context-aware rate limiter.
- **17–20 — applied scenarios:** HTTP middleware, graceful shutdown, parallel file analysis, and a job runner with timeouts, retries, and panic recovery.

Run this track's tests with the race detector, for example: `go test ./tasks/11-worker-pool -race -v`. In an interview, you should be able not only to write the code, but also to explain who closes each channel, how every goroutine terminates, and what happens when the context is canceled.

## Learning Path

| # | Topic | Estimate | Core Practice |
|---:|---|---:|---|
| 01 | [Slices](tasks/01-slices) | 10 min | `append`, `map`, ordering, `nil` |
| 02 | [Unicode strings](tasks/02-unicode) | 15 min | runes instead of bytes |
| 03 | [Errors](tasks/03-errors) | 15 min | wrapping, `errors.Is/As` |
| 04 | [Streaming I/O](tasks/04-stream-io) | 20 min | `io.Reader`, bounded memory |
| 05 | [JSON configuration](tasks/05-json-config) | 25 min | strict decoding and validation |
| 06 | [Atomic file writes](tasks/06-atomic-file) | 25 min | temporary file, rename, cleanup |
| 07 | [File traversal](tasks/07-duplicate-files) | 30 min | `fs.WalkDir`, hashing |
| 08 | [CSV aggregation](tasks/08-csv-aggregation) | 25 min | streaming data parsing |
| 09 | [Generics](tasks/09-generic-set) | 20 min | generic `Set[T]` |
| 10 | [HTTP client](tasks/10-http-client) | 30 min | context, status codes, limits |
| 11 | [Worker pool](tasks/11-worker-pool) | 30 min | goroutines and channels |
| 12 | [Pipeline cancellation](tasks/12-cancellable-pipeline) | 30 min | preventing goroutine leaks |
| 13 | [Fan-out with errors](tasks/13-first-success) | 35 min | early cancellation, `errors.Join` |
| 14 | [TTL cache](tasks/14-ttl-cache) | 40 min | mutexes and clock injection |
| 15 | [Request deduplication](tasks/15-singleflight) | 45 min | single-flight without libraries |
| 16 | [Rate limiter](tasks/16-rate-limiter) | 35 min | time, context, synchronization |
| 17 | [HTTP middleware](tasks/17-http-middleware) | 30 min | handlers, recovery, request IDs |
| 18 | [Graceful shutdown](tasks/18-graceful-shutdown) | 45 min | service lifecycle |
| 19 | [Parallel log analysis](tasks/19-log-analyzer) | 45 min | bounded concurrency, files |
| 20 | [Mini job runner](tasks/20-job-runner) | 60 min | combining all the techniques |

## What May Feel Unfamiliar After Python

- Do not try to translate list comprehensions literally: a plain loop is often the clearest option.
- Decide whether the distinction between a `nil` slice and an empty slice matters.
- An error is an ordinary value. Add context with `%w`; do not compare error text.
- Accept narrow interfaces such as `io.Reader` and return concrete types.
- The code that creates a channel is usually responsible for closing it. Do not close an input channel from the receiver side.
- Every goroutine you start must have a clear termination path.
- Check concurrent code with `-race`; an ordinary passing test does not rule out data races.

## Course Rules

Do not change public signatures unless necessary: the tests define the contract. You may add private functions and types. Premature optimization is not required, but you must follow the memory and concurrency constraints in each task.

Solutions are intentionally omitted from `main`. The `my-solution` branch preserves the author's personal progress and is not a complete answer key. As you work through the course, it is useful to save each step in a separate commit and briefly note in the exercise README which approach you chose and why.
