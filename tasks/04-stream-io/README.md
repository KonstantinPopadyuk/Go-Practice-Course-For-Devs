# 04. Streaming Line Count

**Time:** 15–20 minutes.

Implement `CountLines(r io.Reader)`. Count lines as a stream without loading the entire source into memory.

## Contract

- A line is a sequence terminated by `\n`; a non-empty trailing sequence without `\n` also counts as a line.
- `"a\nb\n"` and `"a\nb"` both contain two lines.
- Empty input contains zero lines; `"\n"` contains one.
- The function must return read errors with additional context and preserve the cause through `%w`.
- Support lines longer than 64 KiB: the default `bufio.Scanner` limit must not cause unexpected failures.

Run: `go test ./tasks/04-stream-io -v`.

Accepting an `io.Reader` lets the same function work with a file, network connection, string, or test double. You can increase the `Scanner` buffer or build the solution on `bufio.Reader`.

## Going Further

Add `CountNonEmptyLines` with a `context.Context` that can be canceled during a slow read.
