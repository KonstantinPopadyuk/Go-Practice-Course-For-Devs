# 02. Unicode-Safe Reverse

**Time:** 10–15 minutes.

Implement `Reverse`, which reverses a string by Unicode code points rather than UTF-8 bytes.

## Contract

- `"hello"` → `"olleh"`.
- `"Привет"` → `"тевирП"`.
- Emoji must not turn into invalid UTF-8.
- An empty string remains empty.
- Invalid UTF-8 sequences may be handled in the standard way through `[]rune` (they become `utf8.RuneError`).

Run: `go test ./tasks/02-unicode -v`.

The main trap after Python is that indexing a string in Go addresses a byte. `len(s)` is the number of bytes, not characters. For this exercise, `[]rune` is sufficient; grapheme clusters such as a family emoji are a separate, more advanced topic.

## Going Further

Write a version that keeps combining graphemes together as a single unit, and explain why the standard library alone is insufficient for this task.
