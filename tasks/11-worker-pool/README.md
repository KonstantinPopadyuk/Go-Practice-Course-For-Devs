# 11. An Order-Preserving Worker Pool

**Time:** 25–30 minutes.

Implement `ParallelMap`: apply a function to every number using a fixed number of worker goroutines.

## Contract

- No more than `workers` calls to `fn` may run concurrently.
- Results must preserve input order, even if the jobs finish in a different order.
- If `workers <= 0`, return `ErrInvalidWorkers`.
- If `fn` returns an error, cancel the remaining work and return an error containing the zero-based index of the element.
- Respect cancellation of the provided context.
- Do not start a goroutine for every element: the number of workers must be bounded.
- Do not leave blocked goroutines behind after the function returns.

Run: `go test ./tasks/11-worker-pool -race -v`.

A practical pattern is to send `{index, value}` through the jobs channel and write each result into a preallocated slice at its unique index. This is safe as long as no two workers write to the same element.

## Going Further

Create a generic version, `ParallelMap[A, B any]`, and compare the readability of the APIs.
