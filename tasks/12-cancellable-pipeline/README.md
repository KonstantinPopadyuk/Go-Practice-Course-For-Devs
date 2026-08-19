# 12. A Cancellable Pipeline

**Time:** 25–30 minutes.

Implement three pipeline stages:

- `Generate(ctx, values...) <-chan int` sends the input numbers;
- `Square(ctx, in) <-chan int` squares the numbers;
- `Take(ctx, in, n) []int` receives at most `n` values.

## Contract

- `Generate` and `Square` each start one goroutine and close only **their own output** channels; `Take` runs synchronously.
- Every blocking send or receive must be able to stop when `ctx.Done()` is closed.
- `Take` returns early if the input channel is closed; if `n <= 0`, it returns an empty slice.
- A typical call such as `Take(ctx, Square(ctx, Generate(ctx, ...)), 3)` must not leak goroutines after the context is canceled.
- Do not buffer the entire input: elements must flow through the pipeline as they become ready.

Run: `go test ./tasks/12-cancellable-pipeline -race -v`.

The key lesson is that an early-exiting consumer stops reading from the channel, so upstream stages must receive a cancellation signal. In the calling code, create a child context and call its cancel function after `Take` returns.

## Going Further

Add `Merge`, which combines several channels and closes its output after all inputs have completed.
